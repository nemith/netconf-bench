# NETCONF Benchmark Suite
# Run with: just <command>

# Default recipe - show available commands
default:
    @just --list

# Setup dependencies and build clients
setup:
    @echo "Setting up NETCONF benchmark suite..."
    @echo "Building server..."
    cd server && go mod tidy && go build -o netconf-server
    @echo "Building Go clients..."
    cd clients/go-nemith && go mod tidy && go build -o benchmark
    cd clients/go-scrapligo && go mod tidy && go build -o benchmark
    @echo "Setting up Python environment..."
    cd clients/ncclient && python3 -m venv .venv && .venv/bin/pip install -q -r requirements.txt
    @echo "Creating results and graphs directories..."
    mkdir -p results graphs
    @echo "Setup complete!"

# Clean build artifacts and results
clean:
    @echo "Cleaning..."
    rm -f server/netconf-server
    rm -f clients/go-nemith/benchmark
    rm -f clients/go-scrapligo/benchmark
    rm -rf clients/ncclient/.venv
    rm -rf results/*.json results/*.csv
    rm -rf graphs/*.png
    @echo "Clean complete!"

# Start NETCONF server in background
server:
    @echo "Starting NETCONF server on port 8830..."
    @./server/netconf-server &
    @echo "Server started. PID: $$!"

# Stop NETCONF server
stop-server:
    @echo "Stopping NETCONF server..."
    @pkill -f netconf-server || true

# Run all benchmarks
bench: stop-server
    @echo "Running full benchmark suite..."
    @./server/netconf-server > /dev/null 2>&1 & echo $$! > /tmp/netconf-server.pid
    @sleep 2
    @just _bench-all
    @kill `cat /tmp/netconf-server.pid` 2>/dev/null || true
    @rm -f /tmp/netconf-server.pid
    @echo "Benchmarks complete! Results in results/"

# Internal: Run all benchmark combinations
_bench-all:
    @echo "Running benchmarks (this may take several minutes)..."
    @just _bench-size-sweep
    @just _bench-count-sweep

# Benchmark response size variations (1KB to 1MB)
_bench-size-sweep:
    #!/usr/bin/env bash
    set -euo pipefail
    echo "=== Response Size Benchmarks (100 requests each) ==="
    for size in 1024 10240 102400 1048576; do
        echo "Testing size: ${size} bytes"
        just _bench-client go-nemith ${size} 100
        just _bench-client go-scrapligo ${size} 100
    done

# Benchmark different request counts (10KB responses)
_bench-count-sweep:
    #!/usr/bin/env bash
    set -euo pipefail
    echo "=== Request Count Benchmarks (10KB responses) ==="
    for count in 50 100 200 500; do
        echo "Testing count: ${count} requests"
        just _bench-client go-nemith 10240 ${count}
        just _bench-client go-scrapligo 10240 ${count}
    done

# Internal: Benchmark a specific client
_bench-client client size count:
    @mkdir -p results
    @if [ "{{client}}" = "go-nemith" ]; then \
        hyperfine --warmup 1 --runs 3 \
            --export-json "results/{{client}}-{{size}}-{{count}}.json" \
            "./clients/go-nemith/benchmark --size {{size}} --count {{count}}"; \
    elif [ "{{client}}" = "go-scrapligo" ]; then \
        hyperfine --warmup 1 --runs 3 \
            --export-json "results/{{client}}-{{size}}-{{count}}.json" \
            "./clients/go-scrapligo/benchmark --size {{size}} --count {{count}}"; \
    elif [ "{{client}}" = "ncclient" ]; then \
        hyperfine --warmup 1 --runs 3 \
            --export-json "results/{{client}}-{{size}}-{{count}}.json" \
            "./clients/ncclient/.venv/bin/python3 ./clients/ncclient/benchmark.py --size {{size}} --count {{count}}"; \
    fi

# Run quick benchmark (fewer iterations for testing)
bench-quick: stop-server
    @echo "Running quick benchmark..."
    @./server/netconf-server > /dev/null 2>&1 & echo $$! > /tmp/netconf-server.pid
    @sleep 2
    @just _bench-client go-nemith 10240 50
    @just _bench-client go-scrapligo 10240 50
    @kill `cat /tmp/netconf-server.pid` 2>/dev/null || true
    @rm -f /tmp/netconf-server.pid
    @echo "Quick benchmark complete!"

# Generate graphs from results
graph:
    @echo "Generating graphs from benchmark results..."
    @python3 scripts/generate_graphs.py
    @echo "Graphs saved to graphs/"

# Run benchmarks and generate graphs
all: bench graph
    @echo "All benchmarks complete with graphs!"

# Test individual client (usage: just test-client go-nemith)
test-client client: stop-server
    @./server/netconf-server > /dev/null 2>&1 & echo $$! > /tmp/netconf-server.pid
    @sleep 2
    @if [ "{{client}}" = "go-nemith" ]; then \
        ./clients/go-nemith/benchmark --size 1024 --count 10 --pipeline 1; \
    elif [ "{{client}}" = "go-scrapligo" ]; then \
        ./clients/go-scrapligo/benchmark --size 1024 --count 10 --pipeline 1; \
    elif [ "{{client}}" = "ncclient" ]; then \
        ./clients/ncclient/.venv/bin/python3 ./clients/ncclient/benchmark.py --size 1024 --count 10 --pipeline 1; \
    fi
    @kill `cat /tmp/netconf-server.pid` 2>/dev/null || true
    @rm -f /tmp/netconf-server.pid

# Show benchmark results summary
results:
    @echo "=== Benchmark Results Summary ==="
    @for f in results/*.json; do \
        if [ -f "$$f" ]; then \
            echo "$$f:"; \
            python3 -c "import json; d=json.load(open('$$f')); print(f\"  Mean: {d['results'][0]['mean']:.3f}s\")"; \
        fi \
    done
