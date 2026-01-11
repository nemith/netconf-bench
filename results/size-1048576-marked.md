| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 638.3 ± 28.0 | 156.8 ± 7.0 | **baseline** | 10 |
| Python: scrapli-netconf (asyncssh) | 203.4 ± 7.0 | 2568.8 ± 45.8 | 38.9 ± 0.7 | 4.02x slower | 10 |
| Python: scrapli-netconf (ssh2) | 92.0 ± 0.7 | 3179.7 ± 409.2 | 32.0 ± 4.4 | 4.98x slower | 10 |
| Python: scrapli-netconf (paramiko) | 130.1 ± 1.4 | 4521.3 ± 93.0 | 22.1 ± 0.5 | 7.08x slower | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 5562.6 ± 168.9 | 18.0 ± 0.5 | 8.71x slower | 10 |
| Python: ncclient (libssh) | 87.9 ± 0.7 | 10547.4 ± 21.8 | 9.5 ± 0.0 | 16.52x slower | 10 |
| Python: ncclient (paramiko) | 144.9 ± 1.3 | 10742.0 ± 79.1 | 9.3 ± 0.1 | 16.83x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 4.1 ± 0.7 | 14642.0 ± 33.4 | 6.8 ± 0.0 | 22.94x slower | 10 |
