| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf)` | 447.3 ± 3.5 | 442.3 | 452.6 | 1.16 ± 0.02 |
| `Go (scrapli/scrapligo)` | 480.0 ± 6.9 | 470.2 | 494.1 | 1.24 ± 0.03 |
| `Python (ncclient/paramiko)` | 10579.5 ± 20.6 | 10560.6 | 10626.7 | 27.38 ± 0.55 |
| `Python (ncclient/libssh)` | 10475.0 ± 6.7 | 10464.1 | 10485.2 | 27.11 ± 0.54 |
| `Python (scrapli-netconf/system)` | 492.5 ± 5.1 | 485.5 | 502.5 | 1.27 ± 0.03 |
| `Python (scrapli-netconf/paramiko)` | 390.0 ± 14.1 | 369.2 | 418.2 | 1.01 ± 0.04 |
| `Python (scrapli-netconf/ssh2)` | 487.4 ± 11.7 | 464.4 | 507.2 | 1.26 ± 0.04 |
| `Python (scrapli-netconf/asyncssh)` | 386.5 ± 7.8 | 378.4 | 400.0 | 1.00 |
