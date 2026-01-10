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
