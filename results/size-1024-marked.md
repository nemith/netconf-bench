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
