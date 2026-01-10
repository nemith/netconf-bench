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
