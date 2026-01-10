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
