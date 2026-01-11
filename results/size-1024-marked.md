| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |
|----------------|------------|----------------|-----|-------------|------|
| Go: nemith.io/netconf (Do) | 0.0 ± 0.0 | 8.0 ± 1.3 | 11768.7 ± 2125.6 | **baseline** | 10 |
| Go: nemith.io/netconf (Exec) | 0.0 ± 0.0 | 17.1 ± 1.6 | 5800.0 ± 591.7 | 2.14x slower | 10 |
| Python: scrapli-netconf (asyncssh) | 200.5 ± 4.6 | 26.0 ± 3.3 | 3890.7 ± 466.1 | 3.25x slower | 10 |
| Go: github.com/scrapli/scrapligo (standard) | 3.7 ± 0.8 | 137.4 ± 3.6 | 724.6 ± 19.0 | 17.18x slower | 10 |
| Python: scrapli-netconf (ssh2) | 92.2 ± 1.6 | 2061.8 ± 4.9 | 48.5 ± 0.1 | 257.73x slower | 10 |
| Python: scrapli-netconf (paramiko) | 130.3 ± 2.2 | 2074.2 ± 5.8 | 48.2 ± 0.1 | 259.27x slower | 10 |
| Python: ncclient (libssh) | 87.8 ± 1.0 | 10078.6 ± 4.7 | 9.9 ± 0.0 | 1259.83x slower | 10 |
| Python: ncclient (paramiko) | 144.4 ± 1.7 | 10083.8 ± 2.4 | 9.9 ± 0.0 | 1260.47x slower | 10 |
