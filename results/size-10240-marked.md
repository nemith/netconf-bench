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
