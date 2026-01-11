| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 72.0 ± 4.3 | 1381.7 ± 81.6 | **baseline** | 10 |
| Python: scrapli-netconf (asyncssh) | 203.0 ± 5.9 | 259.8 ± 11.4 | 385.4 ± 16.4 | 3.61x slower | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 562.6 ± 16.7 | 177.7 ± 5.2 | 7.81x slower | 10 |
| Python: scrapli-netconf (ssh2) | 92.4 ± 0.8 | 999.2 ± 6.2 | 100.1 ± 0.6 | 13.88x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 4.1 ± 0.7 | 1516.0 ± 7.5 | 65.9 ± 0.3 | 21.06x slower | 10 |
| Python: scrapli-netconf (paramiko) | 129.4 ± 1.0 | 1861.2 ± 49.5 | 53.8 ± 1.4 | 25.85x slower | 10 |
| Python: ncclient (libssh) | 88.3 ± 0.8 | 10203.1 ± 9.8 | 9.8 ± 0.0 | 141.71x slower | 10 |
| Python: ncclient (paramiko) | 145.1 ± 1.7 | 10222.8 ± 5.9 | 9.8 ± 0.0 | 141.98x slower | 10 |
