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
