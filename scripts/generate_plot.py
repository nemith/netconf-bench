#!/usr/bin/env -S uv run
# /// script
# requires-python = ">=3.10"
# dependencies = [
#     "matplotlib",
#     "pandas",
# ]
# ///

import sys
import pandas as pd
import matplotlib.pyplot as plt

if len(sys.argv) < 4:
    print(f"Usage: {sys.argv[0]} <output.png> <title> <csv_file1> [csv_file2] ...")
    sys.exit(1)

output_file = sys.argv[1]
title = sys.argv[2]
csv_files = sys.argv[3:]

# Read all CSV files and combine
dfs = []
for csv_file in csv_files:
    df = pd.read_csv(csv_file)
    dfs.append(df)

df = pd.concat(dfs, ignore_index=True)

# Calculate means per implementation and framing
stats = df.groupby(["implementation", "framing"])["rps"].mean().reset_index()

# Create plot
fig, ax = plt.subplots(figsize=(12, 6))

framings = stats["framing"].unique()
implementations = stats["implementation"].unique()
x = range(len(framings))
width = 0.8 / len(implementations)

for i, impl in enumerate(implementations):
    impl_data = stats[stats["implementation"] == impl]
    values = [
        impl_data[impl_data["framing"] == f]["rps"].values[0]
        if len(impl_data[impl_data["framing"] == f]) > 0
        else 0
        for f in framings
    ]
    ax.bar([pos + i * width for pos in x], values, width, label=impl)

ax.set_xlabel("Framing Mode")
ax.set_ylabel("Request Per Second (RPS)")
ax.set_title(title)
ax.set_xticks([pos + width * (len(implementations) - 1) / 2 for pos in x])
ax.set_xticklabels(framings)
ax.legend(bbox_to_anchor=(1.05, 1), loc="upper left")
ax.grid(True, alpha=0.3, axis="y")

plt.tight_layout()
plt.savefig(output_file, dpi=300, bbox_inches="tight")
print(f"Plot saved to {output_file}")
