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
