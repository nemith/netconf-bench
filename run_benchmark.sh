#!/usr/bin/env bash
set -euo pipefail

# --- Configuration ---
RESULTS_DIR="results"
BUILD_DIR="build-out"
SCRIPT_DIR="scripts"

# --- Dependencies Check ---
if ! command -v uv &> /dev/null; then
    echo "Error: 'uv' is required to run the generic plotter."
    echo "Install it: curl -LsSf [https://astral.sh/uv/install.sh](https://astral.sh/uv/install.sh) | sh"
    exit 1
fi

generate_readme() {
    cat > README.md << 'EOF'
# NETCONF Client Benchmark Suite

A comprehensive benchmark suite for comparing NETCONF client implementations.

## Tested Clients

- **Go**: [nemith.io/netconf](https://github.com/nemith/netconf) ✅
- **Go**: [scrapli/scrapligo](https://github.com/scrapli/scrapligo) ✅
- **Python**: [ncclient](https://github.com/ncclient/ncclient) ✅

## Benchmark Results

All tests run 100 requests comparing NETCONF 1.1 (Chunked Framing) vs NETCONF 1.0 (Marked Framing).

### 1KB Responses

![1KB Graph](results/benchmark_1024.png)

#### NETCONF 1.1 (Chunked)
EOF
    cat "${RESULTS_DIR}/size-1024-chunked.md" >> README.md

    cat >> README.md << 'EOF'

#### NETCONF 1.0 (Marked)
EOF
    cat "${RESULTS_DIR}/size-1024-marked.md" >> README.md

    cat >> README.md << 'EOF'

### 10KB Responses

![10KB Graph](results/benchmark_10240.png)

#### NETCONF 1.1 (Chunked)
EOF
    cat "${RESULTS_DIR}/size-10240-chunked.md" >> README.md

    cat >> README.md << 'EOF'

#### NETCONF 1.0 (Marked)
EOF
    cat "${RESULTS_DIR}/size-10240-marked.md" >> README.md

    cat >> README.md << 'EOF'

### 100KB Responses

![100KB Graph](results/benchmark_102400.png)

#### NETCONF 1.1 (Chunked)
EOF
    cat "${RESULTS_DIR}/size-102400-chunked.md" >> README.md

    cat >> README.md << 'EOF'

#### NETCONF 1.0 (Marked)
EOF
    cat "${RESULTS_DIR}/size-102400-marked.md" >> README.md

    cat >> README.md << 'EOF'

### 1MB Responses

![1MB Graph](results/benchmark_1048576.png)

#### NETCONF 1.1 (Chunked)
EOF
    cat "${RESULTS_DIR}/size-1048576-chunked.md" >> README.md

    cat >> README.md << 'EOF'

#### NETCONF 1.0 (Marked)
EOF
    cat "${RESULTS_DIR}/size-1048576-marked.md" >> README.md

    cat >> README.md << 'EOF'

## Running Benchmarks

    ./run_benchmark.sh

## Prerequisites

- Go 1.25+
- [hyperfine](https://github.com/sharkdp/hyperfine)
- [uv](https://github.com/astral-sh/uv) (for python dependency management)

    # macOS
    brew install hyperfine
    curl -LsSf [https://astral.sh/uv/install.sh](https://astral.sh/uv/install.sh) | sh

EOF
}

echo "=== NETCONF Client Benchmark Suite ==="
echo ""

# Build everything
echo "Building server and clients..."
mkdir -p "$BUILD_DIR"
mkdir -p "$RESULTS_DIR"

if [ ! -f "${SCRIPT_DIR}/plot.py" ]; then
    echo "Error: ${SCRIPT_DIR}/plot.py not found!"
    exit 1
fi

echo "  - Building NETCONF server..."
cd server && go build -o "../$BUILD_DIR/netconf-server" && cd ..

echo "  - Building nemith.io/netconf client..."
cd clients/go-nemith && go build -o "../../$BUILD_DIR/nemith-benchmark" && cd ../..

echo "  - Building scrapli/scrapligo client..."
cd clients/go-scrapligo && go build -o "../../$BUILD_DIR/scrapligo-benchmark" && cd ../..

echo "  - Setting up ncclient (Python) environment..."
cd clients/ncclient && uv sync && cd ../..

echo "Build complete!"
echo ""

# --- Define Common Aliases ---
# These arguments clean up the legend in the graph.
PLOT_ALIASES=(
    --alias ".*nemith-benchmark.*" "Go (nemith.io)"
    --alias ".*scrapligo-benchmark.*" "Go (scrapli)"
    --alias ".*ncclient.*" "Python (ncclient)"
)

# --- Size x Framing Mode Benchmarks ---
echo "=== Benchmarking Response Sizes with Framing Modes (100 requests each) ==="
echo ""

# Test each size with both framing modes
for size in 1024 10240 102400 1048576; do
    # Calculate Human Readable Label
    if [ $size -ge 1048576 ]; then label="$((size / 1048576))MB"
    elif [ $size -ge 1024 ]; then label="$((size / 1024))KB"
    else label="${size}B"; fi

    SIZE_ARGS=() # Array to hold arguments for this size's graph

    for framing_mode in "chunked" "marked"; do
        echo "Starting NETCONF server with ${framing_mode} framing..."
        pkill -f netconf-server || true

        # Start server with appropriate framing mode
        if [ "$framing_mode" = "chunked" ]; then
            ./"$BUILD_DIR"/netconf-server --chunked=true > /dev/null 2>&1 &
        else
            ./"$BUILD_DIR"/netconf-server --chunked=false > /dev/null 2>&1 &
        fi
        SERVER_PID=$!

        # Give server time to start
        sleep 1

        echo "Testing ${label} responses with ${framing_mode} framing..."

        JSON_OUT="${RESULTS_DIR}/size-${size}-${framing_mode}.json"

        # Display name for framing mode
        if [ "$framing_mode" = "chunked" ]; then
            framing_label="NETCONF 1.1 (Chunked)"
        else
            framing_label="NETCONF 1.0 (Marked)"
        fi

        hyperfine --warmup 2 \
            --export-markdown "${RESULTS_DIR}/size-${size}-${framing_mode}.md" \
            --export-json "$JSON_OUT" \
            --command-name "Go (nemith.io/netconf)" "./$BUILD_DIR/nemith-benchmark --size ${size} --count 100" \
            --command-name "Go (scrapli/scrapligo)" "./$BUILD_DIR/scrapligo-benchmark --size ${size} --count 100" \
            --command-name "Python (ncclient)" "cd clients/ncclient && ./.venv/bin/python benchmark.py --size ${size} --count 100 && cd ../.."

        echo ""
        kill $SERVER_PID 2>/dev/null || true

        # Append to plotter arguments for this size
        SIZE_ARGS+=(--data "$JSON_OUT" "$framing_label")
    done

    echo "Generating ${label} Graph..."
    uv run "${SCRIPT_DIR}/plot.py" \
        --output "${RESULTS_DIR}/benchmark_${size}.png" \
        --title "Performance with ${label} Responses (100 requests)" \
        --xlabel "Framing Mode" \
        "${PLOT_ALIASES[@]}" \
        "${SIZE_ARGS[@]}"
done

echo "=== Generating README with Results ==="
generate_readme

echo ""
echo "=== Benchmark Complete! ==="
echo "Results written to README.md"
