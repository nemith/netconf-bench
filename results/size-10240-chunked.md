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
