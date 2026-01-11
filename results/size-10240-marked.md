| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 15.2 ± 3.6 | 6672.7 ± 1819.5 | **baseline** | 10 |
| Python: scrapli-netconf (asyncssh) | 197.7 ± 2.9 | 43.9 ± 4.0 | 2297.6 ± 197.0 | 2.89x slower | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 82.2 ± 5.6 | 1215.3 ± 87.0 | 5.41x slower | 10 |
| Python: scrapli-netconf (paramiko) | 129.4 ± 1.4 | 241.5 ± 4.0 | 414.1 ± 6.4 | 15.89x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 4.4 ± 0.8 | 249.4 ± 3.5 | 400.2 ± 5.9 | 16.41x slower | 10 |
| Python: scrapli-netconf (ssh2) | 91.3 ± 0.7 | 2081.0 ± 7.0 | 48.1 ± 0.2 | 136.91x slower | 10 |
| Python: ncclient (libssh) | 87.7 ± 0.8 | 10118.3 ± 3.9 | 9.9 ± 0.0 | 665.68x slower | 10 |
| Python: ncclient (paramiko) | 148.6 ± 1.8 | 10119.2 ± 4.6 | 9.9 ± 0.0 | 665.74x slower | 10 |
