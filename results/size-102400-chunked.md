| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `Go (nemith.io/netconf/exec)` | 416.1 ± 6.1 | 407.6 | 426.2 | 7.87 ± 2.90 |
| `Go (nemith.io/netconf/do)` | 52.8 ± 19.4 | 46.0 | 150.0 | 1.00 |
| `Go (scrapli/scrapligo)` | 1176.1 ± 6.5 | 1164.3 | 1184.5 | 22.26 ± 8.18 |
| `Python (ncclient/paramiko)` | 10651.6 ± 10.8 | 10639.9 | 10676.4 | 201.57 ± 74.10 |
| `Python (ncclient/libssh)` | 10556.8 ± 5.4 | 10550.3 | 10564.2 | 199.77 ± 73.44 |
| `Python (scrapli-netconf/system)` | 30342.3 ± 10.4 | 30328.2 | 30365.7 | 574.19 ± 211.08 |
| `Python (scrapli-netconf/paramiko)` | 374.2 ± 12.5 | 364.2 | 407.3 | 7.08 ± 2.61 |
| `Python (scrapli-netconf/ssh2)` | 484.7 ± 7.3 | 471.1 | 495.4 | 9.17 ± 3.37 |
| `Python (scrapli-netconf/asyncssh)` | 368.4 ± 5.1 | 355.3 | 372.8 | 6.97 ± 2.56 |
