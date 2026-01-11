| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 32.4 ± 2.4 | 3056.1 ± 227.5 | **baseline** | 10 |
| Python: scrapli-netconf (asyncssh) | 197.8 ± 7.5 | 246.2 ± 6.7 | 406.2 ± 10.8 | 7.60x slower | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 545.5 ± 9.3 | 183.2 ± 3.1 | 16.84x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 4.4 ± 0.8 | 1906.3 ± 35.7 | 52.5 ± 1.0 | 58.84x slower | 10 |
| Python: scrapli-netconf (paramiko) | 128.6 ± 1.2 | 4186.8 ± 4.7 | 23.9 ± 0.0 | 129.22x slower | 10 |
| Python: scrapli-netconf (ssh2) | 91.9 ± 0.7 | 4200.7 ± 4.4 | 23.8 ± 0.0 | 129.65x slower | 10 |
| Python: ncclient (libssh) | 87.8 ± 1.0 | 10216.1 ± 10.4 | 9.8 ± 0.0 | 315.31x slower | 10 |
| Python: ncclient (paramiko) | 146.8 ± 2.1 | 10241.7 ± 9.6 | 9.8 ± 0.0 | 316.10x slower | 10 |
