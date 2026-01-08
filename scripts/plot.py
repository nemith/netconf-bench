#!/usr/bin/env -S uv run
# /// script
# requires-python = ">=3.10"
# dependencies = [
#     "matplotlib",
#     "numpy",
# ]
# ///

import argparse
import json
import re
import sys

import matplotlib.pyplot as plt
import numpy as np


def parse_args():
    parser = argparse.ArgumentParser(description="Generic Hyperfine JSON Plotter")

    # Input Data: Expects pairs of "file path" and "label"
    # Example: --data results/1kb.json "1KB" --data results/10kb.json "10KB"
    parser.add_argument(
        "--data",
        nargs=2,
        action="append",
        required=True,
        metavar=("JSON_FILE", "LABEL"),
        help="Path to JSON result and its X-axis label",
    )

    # Cleaning Legend Names: Pairs of "regex" and "replacement"
    parser.add_argument(
        "--alias",
        nargs=2,
        action="append",
        metavar=("REGEX", "REPLACEMENT"),
        help="Regex pattern to clean up command names in legend",
    )

    # Chart Configuration
    parser.add_argument("--title", required=True, help="Chart Title")
    parser.add_argument("--xlabel", required=True, help="X Axis Label")
    parser.add_argument("--ylabel", default="Mean Time (s)", help="Y Axis Label")
    parser.add_argument("--output", required=True, help="Output PNG filename")

    return parser.parse_args()


def clean_command_name(cmd, aliases):
    """Applies regex replacements to clean up the command string."""
    if not aliases:
        return cmd

    name = cmd
    for pattern, replacement in aliases:
        # If the pattern matches, replace it
        if re.search(pattern, name):
            name = re.sub(pattern, replacement, name)
            # We usually stop after first match to prevent chains,
            # or continue if you want multiple passes.
            # Here we return immediately on match for simplicity (Mapping style).
            return name
    return name


def main():
    args = parse_args()

    categories = []  # X-axis labels (e.g., "1KB", "10KB")
    client_data = {}  # {"Client A": [0.1, 0.5], "Client B": [0.2, 0.6]}

    # 1. Process all input files
    for filepath, label in args.data:
        categories.append(label)

        try:
            with open(filepath, "r") as f:
                data = json.load(f)
        except Exception as e:
            print(f"Error reading {filepath}: {e}")
            sys.exit(1)

        # 2. Extract results from this file
        for res in data.get("results", []):
            raw_cmd = res.get("command", "unknown")

            # Clean up the name
            name = clean_command_name(raw_cmd, args.alias)

            if name not in client_data:
                client_data[name] = []

            # Hyperfine JSON stores 'mean' in seconds
            client_data[name].append(res.get("mean", 0))

    # 3. Setup Plot
    if not categories or not client_data:
        print("No data found to plot.")
        sys.exit(1)

    x = np.arange(len(categories))
    width = 0.8 / len(client_data)  # Dynamic bar width based on number of clients

    fig, ax = plt.subplots(figsize=(10, 6))

    # 4. Draw Bars
    multiplier = 0
    # Sort clients alphabetically for consistent legend order
    for client in sorted(client_data.keys()):
        times = client_data[client]

        # Handle missing data if one JSON has fewer clients than another
        # (Append 0s if length mismatches, though hyperfine usually runs all)
        while len(times) < len(categories):
            times.append(0)

        offset = width * multiplier
        # Center the group of bars around the tick
        rects = ax.bar(
            x + offset - (width * len(client_data) / 2) + (width / 2),
            times,
            width,
            label=client,
        )
        multiplier += 1

    # 5. Styling
    ax.set_title(args.title, fontsize=14, pad=20)
    ax.set_ylabel(args.ylabel)
    ax.set_xlabel(args.xlabel)
    ax.set_xticks(x)
    ax.set_xticklabels(categories)
    ax.legend(loc="upper left", bbox_to_anchor=(1, 1))  # Legend outside if crowded
    ax.grid(axis="y", linestyle="--", alpha=0.7)

    plt.tight_layout()
    plt.savefig(args.output, dpi=100)
    print(f"Generated graph: {args.output}")


if __name__ == "__main__":
    main()
