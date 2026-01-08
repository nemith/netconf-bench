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
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 23.8 ± 2.6 | 19.9 | 31.2 | 1.00 |
| `Go (scrapli/scrapligo)` | 117.9 ± 5.7 | 104.8 | 133.1 | 4.96 ± 0.59 |
| `Python (ncclient)` | 10596.0 ± 34.8 | 10537.8 | 10650.2 | 445.77 ± 48.41 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 24.0 ± 2.7 | 19.9 | 37.6 | 1.00 |
| `Go (scrapli/scrapligo)` | 98.2 ± 27.4 | 44.1 | 117.7 | 4.09 ± 1.23 |
| `Python (ncclient)` | 10552.6 ± 25.0 | 10504.8 | 10595.9 | 439.15 ± 50.24 |

### 10KB Responses

![10KB Graph](results/benchmark_10240.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 67.7 ± 4.2 | 58.2 | 79.9 | 1.00 |
| `Go (scrapli/scrapligo)` | 164.3 ± 5.3 | 154.5 | 175.5 | 2.43 ± 0.17 |
| `Python (ncclient)` | 10661.6 ± 63.1 | 10574.8 | 10770.7 | 157.40 ± 9.82 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 67.5 ± 2.9 | 61.0 | 80.7 | 1.00 |
| `Go (scrapli/scrapligo)` | 134.7 ± 20.4 | 76.7 | 153.1 | 1.99 ± 0.31 |
| `Python (ncclient)` | 10627.9 ± 41.6 | 10556.2 | 10700.8 | 157.40 ± 6.81 |

### 100KB Responses

![100KB Graph](results/benchmark_102400.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 388.8 ± 4.2 | 383.1 | 394.4 | 1.00 |
| `Go (scrapli/scrapligo)` | 1195.8 ± 4.3 | 1190.3 | 1204.2 | 3.08 ± 0.03 |
| `Python (ncclient)` | 11175.8 ± 42.3 | 11113.1 | 11266.0 | 28.74 ± 0.33 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 415.6 ± 15.0 | 402.6 | 446.3 | 1.00 |
| `Go (scrapli/scrapligo)` | 474.3 ± 7.4 | 464.3 | 491.0 | 1.14 ± 0.04 |
| `Python (ncclient)` | 10833.0 ± 82.7 | 10678.3 | 10921.2 | 26.07 ± 0.96 |

### 1MB Responses

![1MB Graph](results/benchmark_1048576.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [s] | Min [s] | Max [s] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 3.707 ± 0.056 | 3.608 | 3.805 | 1.00 |
| `Go (scrapli/scrapligo)` | 58.516 ± 0.347 | 58.132 | 59.313 | 15.79 ± 0.26 |
| `Python (ncclient)` | 13.160 ± 0.439 | 12.262 | 13.460 | 3.55 ± 0.13 |

#### NETCONF 1.0 (Marked)
| Command | Mean [s] | Min [s] | Max [s] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 4.506 ± 0.097 | 4.327 | 4.607 | 1.15 ± 0.03 |
| `Go (scrapli/scrapligo)` | 3.912 ± 0.011 | 3.902 | 3.938 | 1.00 |
| `Python (ncclient)` | 12.385 ± 0.122 | 12.104 | 12.494 | 3.17 ± 0.03 |

## Running Benchmarks

    ./run_benchmark.sh

## Prerequisites

- Go 1.25+
- [hyperfine](https://github.com/sharkdp/hyperfine)
- [uv](https://github.com/astral-sh/uv) (for python dependency management)

    # macOS
    brew install hyperfine
    curl -LsSf [https://astral.sh/uv/install.sh](https://astral.sh/uv/install.sh) | sh

