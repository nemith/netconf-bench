| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 29.3 ± 3.6 | 17.9 | 35.8 | 1.24 ± 0.22 |
| `Go (nemith.io/netconf/do)` | 23.8 ± 2.9 | 17.6 | 29.4 | 1.00 |
| `Go (scrapli/scrapligo)` | 119.1 ± 9.6 | 111.8 | 161.1 | 5.01 ± 0.74 |
| `Python (ncclient/paramiko)` | 10427.3 ± 4.8 | 10421.9 | 10435.3 | 438.83 ± 54.46 |
| `Python (ncclient/libssh)` | 10340.7 ± 3.7 | 10336.2 | 10349.6 | 435.19 ± 54.00 |
| `Python (scrapli-netconf/system)` | 30325.9 ± 7.6 | 30308.1 | 30334.6 | 1276.27 ± 158.37 |
| `Python (scrapli-netconf/paramiko)` | 174.7 ± 7.1 | 169.0 | 191.8 | 7.35 ± 0.96 |
| `Python (scrapli-netconf/ssh2)` | 214.3 ± 7.5 | 200.7 | 226.4 | 9.02 ± 1.16 |
| `Python (scrapli-netconf/asyncssh)` | 190.5 ± 6.1 | 178.9 | 199.6 | 8.02 ± 1.03 |
