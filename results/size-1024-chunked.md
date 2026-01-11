| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 7.8 ± 1.7 | 12652.4 ± 3333.2 | **baseline** | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 17.3 ± 1.8 | 5706.0 ± 625.4 | 2.22x slower | 10 |
| Python: scrapli-netconf (asyncssh) | 201.4 ± 3.8 | 30.3 ± 2.7 | 3336.4 ± 282.6 | 3.88x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 4.2 ± 0.8 | 134.1 ± 3.8 | 744.0 ± 20.9 | 17.19x slower | 10 |
| Python: scrapli-netconf (paramiko) | 129.7 ± 1.5 | 4119.2 ± 5.1 | 24.3 ± 0.0 | 528.10x slower | 10 |
| Python: scrapli-netconf (ssh2) | 91.8 ± 0.8 | 4119.8 ± 4.0 | 24.3 ± 0.0 | 528.18x slower | 10 |
| Python: ncclient (paramiko) | 148.4 ± 2.0 | 10080.0 ± 2.7 | 9.9 ± 0.0 | 1292.31x slower | 10 |
| Python: ncclient (libssh) | 88.5 ± 1.3 | 10080.8 ± 2.7 | 9.9 ± 0.0 | 1292.41x slower | 10 |
