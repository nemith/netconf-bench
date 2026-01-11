#!/usr/bin/env -S uv run
# /// script
# requires-python = ">=3.10"
# dependencies = [
#     "pandas",
# ]
# ///
import sys
import pandas as pd

if len(sys.argv) != 4:
    print(f"Usage: {sys.argv[0]} <csv_file> <md_file> <runs>")
    sys.exit(1)

csv_file = sys.argv[1]
md_file = sys.argv[2]
runs = sys.argv[3]

df = pd.read_csv(csv_file)

stats = (
    df.groupby("implementation")
    .agg(
        {
            "setup_ms": ["mean", "std"],
            "rpc_calls_ms": ["mean", "std"],
            "rps": ["mean", "std"],
        }
    )
    .round(3)
)

stats.columns = ["_".join(col).strip() for col in stats.columns.values]
stats = stats.sort_values("rpc_calls_ms_mean")

# Find baseline (fastest)
baseline_time = stats["rpc_calls_ms_mean"].iloc[0]
baseline_name = stats.index[0]

with open(md_file, "w") as f:
    f.write(
        "| Implementation | Setup (ms) | RPC Calls (ms) | RPS | vs Baseline | Runs |\n"
    )
    f.write(
        "|----------------|------------|----------------|-----|-------------|------|\n"
    )

    for impl in stats.index:
        setup_mean = stats.loc[impl, "setup_ms_mean"]
        setup_std = stats.loc[impl, "setup_ms_std"]
        rpc_mean = stats.loc[impl, "rpc_calls_ms_mean"]
        rpc_std = stats.loc[impl, "rpc_calls_ms_std"]
        rps_mean = stats.loc[impl, "rps_mean"]
        rps_std = stats.loc[impl, "rps_std"]

        # Calculate slowdown factor
        if impl == baseline_name:
            speedup = "**baseline**"
        else:
            factor = rpc_mean / baseline_time
            speedup = f"{factor:.2f}x slower"

        f.write(
            f"| {impl} | {setup_mean:.1f} ± {setup_std:.1f} | {rpc_mean:.1f} ± {rpc_std:.1f} | {rps_mean:.1f} ± {rps_std:.1f} | {speedup} | {runs} |\n"
        )
