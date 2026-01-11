| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 266.8 ± 8.1 | 374.3 ± 11.0 | **baseline** | 10 |
| Python: scrapli-netconf (asyncssh) | 200.9 ± 4.2 | 2374.4 ± 38.4 | 42.1 ± 0.7 | 8.90x slower | 10 |
| Python: scrapli-netconf (paramiko) | 128.6 ± 1.9 | 4482.1 ± 93.4 | 22.3 ± 0.5 | 16.80x slower | 10 |
| Python: scrapli-netconf (ssh2) | 92.1 ± 1.2 | 4788.9 ± 60.4 | 20.9 ± 0.3 | 17.95x slower | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 5211.9 ± 145.2 | 19.2 ± 0.5 | 19.53x slower | 10 |
| Python: ncclient (libssh) | 87.4 ± 0.7 | 11332.0 ± 28.9 | 8.8 ± 0.0 | 42.47x slower | 10 |
| Python: ncclient (paramiko) | 147.6 ± 3.2 | 11747.4 ± 58.5 | 8.5 ± 0.0 | 44.03x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 4.7 ± 0.5 | 71244.9 ± 234.4 | 1.4 ± 0.0 | 267.03x slower | 10 |
