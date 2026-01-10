# NETCONF Client Benchmark Suite

A comprehensive benchmark suite for comparing NETCONF client implementations.

## Tested Clients

- **Go**: [nemith.io/netconf](https://github.com/nemith/netconf) (Exec) ✅
- **Go**: [nemith.io/netconf](https://github.com/nemith/netconf) (Do) ✅
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
| `Go (nemith.io/netconf/exec)` | 29.3 ± 3.6 | 17.9 | 35.8 | 1.24 ± 0.22 |
| `Go (nemith.io/netconf/do)` | 23.8 ± 2.9 | 17.6 | 29.4 | 1.00 |
| `Go (scrapli/scrapligo)` | 119.1 ± 9.6 | 111.8 | 161.1 | 5.01 ± 0.74 |
| `Python (ncclient/paramiko)` | 10427.3 ± 4.8 | 10421.9 | 10435.3 | 438.83 ± 54.46 |
| `Python (ncclient/libssh)` | 10340.7 ± 3.7 | 10336.2 | 10349.6 | 435.19 ± 54.00 |
| `Python (scrapli-netconf/system)` | 30325.9 ± 7.6 | 30308.1 | 30334.6 | 1276.27 ± 158.37 |
| `Python (scrapli-netconf/paramiko)` | 174.7 ± 7.1 | 169.0 | 191.8 | 7.35 ± 0.96 |
| `Python (scrapli-netconf/ssh2)` | 214.3 ± 7.5 | 200.7 | 226.4 | 9.02 ± 1.16 |
| `Python (scrapli-netconf/asyncssh)` | 190.5 ± 6.1 | 178.9 | 199.6 | 8.02 ± 1.03 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 30.4 ± 2.3 | 22.8 | 35.3 | 1.35 ± 0.20 |
| `Go (nemith.io/netconf/do)` | 22.6 ± 2.8 | 16.7 | 28.0 | 1.00 |
| `Go (scrapli/scrapligo)` | 108.0 ± 18.9 | 44.5 | 119.2 | 4.78 ± 1.03 |
| `Python (ncclient/paramiko)` | 10428.4 ± 7.5 | 10422.5 | 10447.6 | 461.56 ± 58.06 |
| `Python (ncclient/libssh)` | 10341.9 ± 6.8 | 10335.8 | 10355.7 | 457.73 ± 57.57 |
| `Python (scrapli-netconf/system)` | 234.3 ± 3.9 | 230.0 | 244.7 | 10.37 ± 1.32 |
| `Python (scrapli-netconf/paramiko)` | 167.2 ± 9.9 | 151.4 | 187.7 | 7.40 ± 1.03 |
| `Python (scrapli-netconf/ssh2)` | 204.5 ± 7.0 | 194.4 | 218.2 | 9.05 ± 1.18 |
| `Python (scrapli-netconf/asyncssh)` | 180.3 ± 6.5 | 166.2 | 191.4 | 7.98 ± 1.04 |

### 10KB Responses

![10KB Graph](results/benchmark_10240.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 69.3 ± 2.3 | 66.0 | 77.4 | 2.58 ± 0.22 |
| `Go (nemith.io/netconf/do)` | 26.9 ± 2.1 | 20.4 | 32.4 | 1.00 |
| `Go (scrapli/scrapligo)` | 166.4 ± 3.3 | 159.8 | 171.6 | 6.18 ± 0.50 |
| `Python (ncclient/paramiko)` | 10454.9 ± 11.9 | 10442.9 | 10483.4 | 388.38 ± 30.47 |
| `Python (ncclient/libssh)` | 10362.9 ± 4.5 | 10357.2 | 10373.5 | 384.96 ± 30.20 |
| `Python (scrapli-netconf/system)` | 30325.2 ± 7.6 | 30310.7 | 30336.9 | 1126.52 ± 88.36 |
| `Python (scrapli-netconf/paramiko)` | 185.7 ± 4.6 | 174.9 | 195.1 | 6.90 ± 0.57 |
| `Python (scrapli-netconf/ssh2)` | 238.6 ± 10.5 | 220.1 | 255.6 | 8.86 ± 0.80 |
| `Python (scrapli-netconf/asyncssh)` | 204.5 ± 5.8 | 198.4 | 217.7 | 7.60 ± 0.63 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 73.9 ± 3.7 | 64.4 | 84.2 | 2.36 ± 0.22 |
| `Go (nemith.io/netconf/do)` | 31.3 ± 2.5 | 23.4 | 36.5 | 1.00 |
| `Go (scrapli/scrapligo)` | 141.1 ± 1.4 | 138.5 | 143.5 | 4.51 ± 0.36 |
| `Python (ncclient/paramiko)` | 10466.9 ± 38.6 | 10436.5 | 10541.9 | 334.81 ± 26.57 |
| `Python (ncclient/libssh)` | 10452.0 ± 2.6 | 10448.4 | 10456.0 | 334.34 ± 26.50 |
| `Python (scrapli-netconf/system)` | 262.1 ± 8.2 | 249.6 | 279.8 | 8.38 ± 0.71 |
| `Python (scrapli-netconf/paramiko)` | 190.1 ± 9.3 | 175.2 | 205.6 | 6.08 ± 0.57 |
| `Python (scrapli-netconf/ssh2)` | 234.6 ± 5.8 | 222.2 | 244.5 | 7.50 ± 0.62 |
| `Python (scrapli-netconf/asyncssh)` | 193.5 ± 3.5 | 190.1 | 203.1 | 6.19 ± 0.50 |

### 100KB Responses

![100KB Graph](results/benchmark_102400.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 416.1 ± 6.1 | 407.6 | 426.2 | 7.87 ± 2.90 |
| `Go (nemith.io/netconf/do)` | 52.8 ± 19.4 | 46.0 | 150.0 | 1.00 |
| `Go (scrapli/scrapligo)` | 1176.1 ± 6.5 | 1164.3 | 1184.5 | 22.26 ± 8.18 |
| `Python (ncclient/paramiko)` | 10651.6 ± 10.8 | 10639.9 | 10676.4 | 201.57 ± 74.10 |
| `Python (ncclient/libssh)` | 10556.8 ± 5.4 | 10550.3 | 10564.2 | 199.77 ± 73.44 |
| `Python (scrapli-netconf/system)` | 30342.3 ± 10.4 | 30328.2 | 30365.7 | 574.19 ± 211.08 |
| `Python (scrapli-netconf/paramiko)` | 374.2 ± 12.5 | 364.2 | 407.3 | 7.08 ± 2.61 |
| `Python (scrapli-netconf/ssh2)` | 484.7 ± 7.3 | 471.1 | 495.4 | 9.17 ± 3.37 |
| `Python (scrapli-netconf/asyncssh)` | 368.4 ± 5.1 | 355.3 | 372.8 | 6.97 ± 2.56 |

#### NETCONF 1.0 (Marked)
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 442.5 ± 4.1 | 435.3 | 447.6 | 5.58 ± 0.09 |
| `Go (nemith.io/netconf/do)` | 79.3 ± 1.1 | 77.5 | 82.6 | 1.00 |
| `Go (scrapli/scrapligo)` | 465.7 ± 5.9 | 459.6 | 481.7 | 5.87 ± 0.11 |
| `Python (ncclient/paramiko)` | 10615.6 ± 13.5 | 10594.8 | 10641.1 | 133.90 ± 1.91 |
| `Python (ncclient/libssh)` | 10553.7 ± 9.8 | 10529.7 | 10567.1 | 133.12 ± 1.90 |
| `Python (scrapli-netconf/system)` | 489.0 ± 7.2 | 474.8 | 498.9 | 6.17 ± 0.13 |
| `Python (scrapli-netconf/paramiko)` | 362.8 ± 9.6 | 340.6 | 374.2 | 4.58 ± 0.14 |
| `Python (scrapli-netconf/ssh2)` | 481.0 ± 9.3 | 467.3 | 499.9 | 6.07 ± 0.15 |
| `Python (scrapli-netconf/asyncssh)` | 370.2 ± 6.2 | 359.2 | 380.1 | 4.67 ± 0.10 |

### 1MB Responses

![1MB Graph](results/benchmark_1048576.png)

#### NETCONF 1.1 (Chunked)
| Command | Mean [s] | Min [s] | Max [s] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 3.816 ± 0.022 | 3.776 | 3.853 | 12.35 ± 0.17 |
| `Go (nemith.io/netconf/do)` | 0.309 ± 0.004 | 0.304 | 0.315 | 1.00 |
| `Go (scrapli/scrapligo)` | 61.088 ± 0.182 | 60.761 | 61.389 | 197.76 ± 2.49 |
| `Python (ncclient/paramiko)` | 12.038 ± 0.045 | 11.948 | 12.078 | 38.97 ± 0.50 |
| `Python (ncclient/libssh)` | 11.898 ± 0.025 | 11.839 | 11.934 | 38.52 ± 0.48 |
| `Python (scrapli-netconf/system)` | 30.351 ± 0.007 | 30.337 | 30.359 | 98.26 ± 1.20 |
| `Python (scrapli-netconf/paramiko)` | 2.145 ± 0.014 | 2.122 | 2.163 | 6.94 ± 0.10 |
| `Python (scrapli-netconf/ssh2)` | 2.803 ± 0.016 | 2.782 | 2.828 | 9.07 ± 0.12 |
| `Python (scrapli-netconf/asyncssh)` | 2.015 ± 0.020 | 1.978 | 2.049 | 6.52 ± 0.10 |

#### NETCONF 1.0 (Marked)
| Command | Mean [s] | Min [s] | Max [s] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 4.614 ± 0.100 | 4.354 | 4.687 | 5.38 ± 0.13 |
| `Go (nemith.io/netconf/do)` | 0.858 ± 0.008 | 0.844 | 0.875 | 1.00 |
| `Go (scrapli/scrapligo)` | 3.934 ± 0.028 | 3.890 | 3.965 | 4.59 ± 0.05 |
| `Python (ncclient/paramiko)` | 11.167 ± 0.132 | 10.971 | 11.278 | 13.02 ± 0.20 |
| `Python (ncclient/libssh)` | 11.081 ± 0.029 | 11.002 | 11.099 | 12.92 ± 0.13 |
| `Python (scrapli-netconf/system)` | 2.625 ± 0.012 | 2.608 | 2.645 | 3.06 ± 0.03 |
| `Python (scrapli-netconf/paramiko)` | 2.165 ± 0.029 | 2.092 | 2.195 | 2.52 ± 0.04 |
| `Python (scrapli-netconf/ssh2)` | 2.840 ± 0.021 | 2.810 | 2.884 | 3.31 ± 0.04 |
| `Python (scrapli-netconf/asyncssh)` | 1.885 ± 0.017 | 1.863 | 1.919 | 2.20 ± 0.03 |

## Running Benchmarks

    ./run_benchmark.sh

## Prerequisites

- Go 1.25+
- [hyperfine](https://github.com/sharkdp/hyperfine)
- [uv](https://github.com/astral-sh/uv) (for python dependency management)

    # macOS
    brew install hyperfine
    curl -LsSf [https://astral.sh/uv/install.sh](https://astral.sh/uv/install.sh) | sh

