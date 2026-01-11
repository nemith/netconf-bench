#!/usr/bin/env bash
set -euo pipefail

# --- Configuration ---
RESULTS_DIR="results"
BUILD_DIR="build-out"
SCRIPT_DIR="scripts"
RUNS=10

# --- Dependencies Check ---
if ! command -v uv &> /dev/null; then
    echo "Error: 'uv' is required to run the generic plotter."
    echo "Install it: curl -LsSf https://astral.sh/uv/install.sh | sh"
    exit 1
fi

generate_readme() {
    cat > README.md << 'EOF'
# NETCONF Client Benchmark Suite

A comprehensive benchmark suite for comparing NETCONF client implementations.

## Tested Clients

- **Go**: [nemith.io/netconf](https://github.com/nemith/netconf) (Session.Exec) 
- **Go**: [nemith.io/netconf](https://github.com/nemith/netconf) (Session.Do) 
- **Go**: [scrapli/scrapligo](https://github.com/scrapli/scrapligo) 
- **Python**: [ncclient](https://github.com/ncclient/ncclient) (paramiko) 
- **Python**: [ncclient](https://github.com/ncclient/ncclient) (libssh) 
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (system) 
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (paramiko) 
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (ssh2) 
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (asyncssh) 

## Interpretation
These benchmarks should be taken with a grain of salt.  This is not intended to be a definitive performance comparison, but rather a way to get a general idea of how different implementations perform under similar conditions. 

This is not intended to compare language performance (Go vs Python), but rather the efficiency of the client libraries and their NETCONF implementations.   

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
- Python 3.8+ with pandas
- [uv](https://github.com/astral-sh/uv) (for python dependency management)

    # macOS
    pip install pandas
    curl -LsSf https://astral.sh/uv/install.sh | sh

EOF
}

run_benchmark() {
    local cmd="$1"
    local csv_file="$2"
    
    echo -n "    Run $run/$RUNS... "
    
    output=$(eval $cmd)
    csv_line=$(echo "$output" | tail -n 1)
    
    echo "$csv_line" | awk -F',' -v run="$run" '{print $1","$2","$3","$4","run","$5","$6","$7}' >> "$csv_file"
    
    echo "done"
}

echo "=== NETCONF Client Benchmark Suite ==="
echo ""

echo "Building server and clients..."
mkdir -p "$BUILD_DIR"
mkdir -p "$RESULTS_DIR"

echo "  - Building NETCONF server..."
cd server && go build -o "../$BUILD_DIR/netconf-server" && cd ..

echo "  - Building nemith.io/netconf client..."
cd clients/go-nemith && go build -o "../../$BUILD_DIR/nemith-benchmark" && cd ../..

echo "  - Building scrapli/scrapligo client..."
cd clients/go-scrapligo && go build -o "../../$BUILD_DIR/scrapligo-benchmark" && cd ../..

echo "  - Setting up ncclient (Python) environment..."
cd clients/ncclient && uv sync && cd ../..

echo "  - Setting up scrapli-netconf (Python) environment..."
cd clients/scrapli-netconf && uv sync && cd ../..

echo "Build complete!"
echo ""

for size in 1024 10240 102400 1048576; do
    if [ $size -ge 1048576 ]; then label="$((size / 1048576))MB"
    elif [ $size -ge 1024 ]; then label="$((size / 1024))KB"
    else label="${size}B"; fi

    for framing_mode in "chunked" "marked"; do
        echo "Starting NETCONF server with ${framing_mode} framing..."
        pkill -f netconf-server || true

        if [ "$framing_mode" = "chunked" ]; then
            ./"$BUILD_DIR"/netconf-server --chunked=true > /dev/null 2>&1 &
        else
            ./"$BUILD_DIR"/netconf-server --chunked=false > /dev/null 2>&1 &
        fi
        SERVER_PID=$!

        sleep 1

        echo "Testing ${label} responses with ${framing_mode} framing..."

        CSV_OUT="${RESULTS_DIR}/size-${size}-${framing_mode}.csv"
        
        echo "implementation,framing,size,count,run,setup_ms,rpc_calls_ms,rps" > "$CSV_OUT"


        declare -A IMPLEMENTATIONS=(
            ["nemith-exec"]="./$BUILD_DIR/nemith-benchmark --framing ${framing_mode} --method exec --size ${size} --count 100"
            ["nemith-do"]="./$BUILD_DIR/nemith-benchmark --framing ${framing_mode} --method do --size ${size} --count 100"
            ["scrapligo"]="./$BUILD_DIR/scrapligo-benchmark --framing ${framing_mode} --size ${size} --count 100"
            ["ncclient-paramiko"]="(cd clients/ncclient && ./.venv/bin/python benchmark.py --framing ${framing_mode} --backend paramiko --size ${size} --count 100)"
            ["ncclient-libssh"]="(cd clients/ncclient && ./.venv/bin/python benchmark.py --framing ${framing_mode} --backend libssh --size ${size} --count 100)"
           # ["scrapli-system"]="(cd clients/scrapli-netconf && ./.venv/bin/python benchmark.py --framing ${framing_mode} --transport system --size ${size} --count 100)"
            ["scrapli-paramiko"]="(cd clients/scrapli-netconf && ./.venv/bin/python benchmark.py --framing ${framing_mode} --transport paramiko --size ${size} --count 100)"
            ["scrapli-ssh2"]="(cd clients/scrapli-netconf && ./.venv/bin/python benchmark.py --framing ${framing_mode} --transport ssh2 --size ${size} --count 100)"
            ["scrapli-asyncssh"]="(cd clients/scrapli-netconf && ./.venv/bin/python benchmark.py --framing ${framing_mode} --transport asyncssh --size ${size} --count 100)"
        )

        for impl_key in "${!IMPLEMENTATIONS[@]}"; do
            echo "  Testing: $impl_key"
            cmd="${IMPLEMENTATIONS[$impl_key]}"
            
            for run in $(seq 1 $RUNS); do
                run_benchmark "$cmd" "$CSV_OUT"
            done
        done

        echo ""
        
        "${SCRIPT_DIR}/generate_markdown.py" "$CSV_OUT" "${RESULTS_DIR}/size-${size}-${framing_mode}.md" "$RUNS"

        kill $SERVER_PID 2>/dev/null || true
    done

    echo "Generating ${label} Graph..."
    "${SCRIPT_DIR}/generate_plot.py" \
        "${RESULTS_DIR}/benchmark_${size}.png" \
        "Performance with ${label} Responses (100 requests, ${RUNS} runs avg)" \
        "${RESULTS_DIR}/size-${size}-chunked.csv" \
        "${RESULTS_DIR}/size-${size}-marked.csv"
done

echo "=== Generating README with Results ==="
generate_readme

echo ""
echo "=== Benchmark Complete! ==="
echo "Results written to README.md"
echo "Raw data available in ${RESULTS_DIR}/*.csv"