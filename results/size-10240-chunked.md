| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 11.7 ± 2.2 | 8453.4 ± 1729.6 | **baseline** | 10 |
| Python: scrapli-netconf (asyncssh) | 199.9 ± 4.9 | 49.5 ± 5.1 | 2031.0 ± 193.0 | 4.23x slower | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 77.3 ± 4.3 | 1288.6 ± 72.5 | 6.61x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 3.9 ± 1.0 | 266.7 ± 5.2 | 374.4 ± 7.4 | 22.79x slower | 10 |
| Python: scrapli-netconf (ssh2) | 92.1 ± 1.0 | 4116.9 ± 3.4 | 24.3 ± 0.0 | 351.87x slower | 10 |
| Python: scrapli-netconf (paramiko) | 130.0 ± 1.6 | 4117.7 ± 2.8 | 24.3 ± 0.0 | 351.94x slower | 10 |
| Python: ncclient (libssh) | 87.7 ± 0.7 | 10120.6 ± 7.5 | 9.9 ± 0.0 | 865.01x slower | 10 |
| Python: ncclient (paramiko) | 144.9 ± 1.9 | 10121.9 ± 4.4 | 9.9 ± 0.0 | 865.12x slower | 10 |
