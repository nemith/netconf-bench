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
