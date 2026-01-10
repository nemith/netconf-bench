# NETCONF Client Benchmark Suite

A comprehensive benchmark suite for comparing NETCONF client implementations.

## Tested Clients

- **Go**: [nemith.io/netconf](https://github.com/nemith/netconf) ✅
- **Go**: [scrapli/scrapligo](https://github.com/scrapli/scrapligo) ✅
- **Python**: [ncclient](https://github.com/ncclient/ncclient) (paramiko) ✅
- **Python**: [ncclient](https://github.com/ncclient/ncclient) (libssh) ✅
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (system) ✅
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (paramiko) ✅
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (ssh2) ✅
- **Python**: [scrapli/scrapli-netconf](https://github.com/scrapli/scrapli_netconf) (asyncssh) ✅

## Benchmark Results

All tests run 100 requests comparing NETCONF 1.1 (Chunked Framing) vs NETCONF 1.0 (Marked Framing).

### 1KB Responses

![1KB Graph](results/benchmark_1024.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 23.1 ± 2.1 | 20.6 | 32.2 | 1.00 |
| `Go (scrapli/scrapligo)` | 114.1 ± 3.8 | 104.1 | 119.6 | 4.93 ± 0.48 |
| `Python (ncclient/paramiko)` | 10433.1 ± 8.0 | 10420.0 | 10447.1 | 451.19 ± 40.74 |
| `Python (ncclient/libssh)` | 10341.7 ± 6.8 | 10332.8 | 10351.7 | 447.24 ± 40.38 |
| `Python (scrapli-netconf/system)` | 30325.3 ± 6.4 | 30314.2 | 30333.8 | 1311.47 ± 118.41 |
| `Python (scrapli-netconf/paramiko)` | 167.3 ± 8.0 | 150.2 | 182.4 | 7.24 ± 0.74 |
| `Python (scrapli-netconf/ssh2)` | 194.3 ± 8.8 | 180.2 | 207.1 | 8.40 ± 0.85 |
| `Python (scrapli-netconf/asyncssh)` | 185.8 ± 11.9 | 170.0 | 209.3 | 8.03 ± 0.89 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 24.3 ± 2.5 | 21.4 | 32.7 | 1.00 |
| `Go (scrapli/scrapligo)` | 101.6 ± 21.9 | 37.4 | 115.8 | 4.19 ± 1.00 |
| `Python (ncclient/paramiko)` | 10426.9 ± 7.9 | 10419.8 | 10445.2 | 429.85 ± 43.72 |
| `Python (ncclient/libssh)` | 10358.9 ± 41.6 | 10325.1 | 10419.7 | 427.05 ± 43.47 |
| `Python (scrapli-netconf/system)` | 231.8 ± 5.9 | 224.8 | 247.4 | 9.56 ± 1.00 |
| `Python (scrapli-netconf/paramiko)` | 165.0 ± 8.8 | 151.5 | 187.8 | 6.80 ± 0.78 |
| `Python (scrapli-netconf/ssh2)` | 189.3 ± 9.2 | 177.5 | 209.5 | 7.80 ± 0.88 |
| `Python (scrapli-netconf/asyncssh)` | 175.6 ± 6.8 | 166.4 | 187.7 | 7.24 ± 0.79 |

### 10KB Responses

![10KB Graph](results/benchmark_10240.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 67.1 ± 3.2 | 58.6 | 74.3 | 1.00 |
| `Go (scrapli/scrapligo)` | 163.4 ± 3.9 | 157.6 | 172.0 | 2.44 ± 0.13 |
| `Python (ncclient/paramiko)` | 10533.9 ± 12.5 | 10516.5 | 10563.9 | 157.01 ± 7.46 |
| `Python (ncclient/libssh)` | 10452.9 ± 6.5 | 10441.0 | 10461.1 | 155.81 ± 7.40 |
| `Python (scrapli-netconf/system)` | 30333.1 ± 9.3 | 30316.9 | 30352.7 | 452.14 ± 21.48 |
| `Python (scrapli-netconf/paramiko)` | 198.8 ± 10.3 | 188.5 | 226.0 | 2.96 ± 0.21 |
| `Python (scrapli-netconf/ssh2)` | 219.1 ± 9.8 | 203.4 | 237.8 | 3.27 ± 0.21 |
| `Python (scrapli-netconf/asyncssh)` | 211.6 ± 10.0 | 193.5 | 230.6 | 3.15 ± 0.21 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 73.1 ± 4.7 | 65.0 | 83.5 | 1.00 |
| `Go (scrapli/scrapligo)` | 130.8 ± 19.0 | 70.9 | 142.2 | 1.79 ± 0.28 |
| `Python (ncclient/paramiko)` | 10468.4 ± 8.9 | 10453.7 | 10481.0 | 143.26 ± 9.14 |
| `Python (ncclient/libssh)` | 10369.4 ± 9.5 | 10352.8 | 10383.6 | 141.91 ± 9.06 |
| `Python (scrapli-netconf/system)` | 265.2 ± 5.3 | 259.3 | 275.1 | 3.63 ± 0.24 |
| `Python (scrapli-netconf/paramiko)` | 194.2 ± 12.5 | 177.8 | 221.3 | 2.66 ± 0.24 |
| `Python (scrapli-netconf/ssh2)` | 230.3 ± 10.7 | 209.1 | 249.3 | 3.15 ± 0.25 |
| `Python (scrapli-netconf/asyncssh)` | 203.6 ± 7.4 | 192.6 | 219.8 | 2.79 ± 0.20 |

### 100KB Responses

![100KB Graph](results/benchmark_102400.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 424.8 ± 7.6 | 414.0 | 439.7 | 1.13 ± 0.03 |
| `Go (scrapli/scrapligo)` | 1175.3 ± 5.0 | 1166.7 | 1185.0 | 3.13 ± 0.08 |
| `Python (ncclient/paramiko)` | 10582.7 ± 8.2 | 10569.4 | 10595.2 | 28.16 ± 0.71 |
| `Python (ncclient/libssh)` | 10480.1 ± 7.0 | 10472.5 | 10495.0 | 27.89 ± 0.70 |
| `Python (scrapli-netconf/system)` | 30337.3 ± 6.9 | 30329.7 | 30355.3 | 80.72 ± 2.04 |
| `Python (scrapli-netconf/paramiko)` | 375.8 ± 9.5 | 357.7 | 391.3 | 1.00 |
| `Python (scrapli-netconf/ssh2)` | 492.1 ± 17.0 | 475.0 | 533.5 | 1.31 ± 0.06 |
| `Python (scrapli-netconf/asyncssh)` | 381.8 ± 6.8 | 364.5 | 388.9 | 1.02 ± 0.03 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 447.3 ± 3.5 | 442.3 | 452.6 | 1.16 ± 0.02 |
| `Go (scrapli/scrapligo)` | 480.0 ± 6.9 | 470.2 | 494.1 | 1.24 ± 0.03 |
| `Python (ncclient/paramiko)` | 10579.5 ± 20.6 | 10560.6 | 10626.7 | 27.38 ± 0.55 |
| `Python (ncclient/libssh)` | 10475.0 ± 6.7 | 10464.1 | 10485.2 | 27.11 ± 0.54 |
| `Python (scrapli-netconf/system)` | 492.5 ± 5.1 | 485.5 | 502.5 | 1.27 ± 0.03 |
| `Python (scrapli-netconf/paramiko)` | 390.0 ± 14.1 | 369.2 | 418.2 | 1.01 ± 0.04 |
| `Python (scrapli-netconf/ssh2)` | 487.4 ± 11.7 | 464.4 | 507.2 | 1.26 ± 0.04 |
| `Python (scrapli-netconf/asyncssh)` | 386.5 ± 7.8 | 378.4 | 400.0 | 1.00 |

### 1MB Responses

![1MB Graph](results/benchmark_1048576.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [s] | Min [s] | Max [s] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 3.944 ± 0.040 | 3.883 | 4.007 | 1.89 ± 0.03 |
| `Go (scrapli/scrapligo)` | 59.781 ± 0.360 | 59.021 | 60.309 | 28.59 ± 0.46 |
| `Python (ncclient/paramiko)` | 11.973 ± 0.026 | 11.932 | 12.016 | 5.73 ± 0.09 |
| `Python (ncclient/libssh)` | 11.800 ± 0.012 | 11.779 | 11.814 | 5.64 ± 0.08 |
| `Python (scrapli-netconf/system)` | 30.345 ± 0.008 | 30.332 | 30.360 | 14.52 ± 0.21 |
| `Python (scrapli-netconf/paramiko)` | 2.208 ± 0.035 | 2.148 | 2.275 | 1.06 ± 0.02 |
| `Python (scrapli-netconf/ssh2)` | 2.932 ± 0.023 | 2.907 | 2.979 | 1.40 ± 0.02 |
| `Python (scrapli-netconf/asyncssh)` | 2.091 ± 0.031 | 2.057 | 2.163 | 1.00 |

#### NETCONF 1.0 (Marked)
| Command | Mean [s] | Min [s] | Max [s] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 4.607 ± 0.036 | 4.554 | 4.665 | 2.24 ± 0.06 |
| `Go (scrapli/scrapligo)` | 3.946 ± 0.012 | 3.928 | 3.969 | 1.92 ± 0.05 |
| `Python (ncclient/paramiko)` | 11.055 ± 0.017 | 11.031 | 11.093 | 5.38 ± 0.14 |
| `Python (ncclient/libssh)` | 10.932 ± 0.010 | 10.918 | 10.948 | 5.32 ± 0.14 |
| `Python (scrapli-netconf/system)` | 2.601 ± 0.047 | 2.548 | 2.677 | 1.27 ± 0.04 |
| `Python (scrapli-netconf/paramiko)` | 2.099 ± 0.065 | 2.034 | 2.211 | 1.02 ± 0.04 |
| `Python (scrapli-netconf/ssh2)` | 2.950 ± 0.044 | 2.869 | 3.010 | 1.44 ± 0.04 |
| `Python (scrapli-netconf/asyncssh)` | 2.054 ± 0.054 | 1.977 | 2.138 | 1.00 |

## Running Benchmarks

    ./run_benchmark.sh

## Prerequisites

- Go 1.25+
- [hyperfine](https://github.com/sharkdp/hyperfine)
- [uv](https://github.com/astral-sh/uv) (for python dependency management)

    # macOS
    brew install hyperfine
    curl -LsSf [https://astral.sh/uv/install.sh](https://astral.sh/uv/install.sh) | sh

