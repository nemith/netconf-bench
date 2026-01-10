#!/usr/bin/env python3

import argparse
import logging
import sys
import time
from concurrent.futures import ThreadPoolExecutor, as_completed

from ncclient import manager
from ncclient.xml_ import to_ele

# Configure app logging
logging.basicConfig(
    level=logging.INFO, format="%(levelname)s:%(message)s", stream=sys.stderr
)

# Silence ncclient's verbose logging
logging.getLogger("ncclient").setLevel(logging.WARNING)

logger = logging.getLogger(__name__)


def create_filter(size):
    """Create a NETCONF filter with size attribute."""
    return f'<filter type="subtree"><size>{size}</size></filter>'


def run_get_request(mgr, size):
    """Execute a single get request."""
    filter_xml = create_filter(size)
    try:
        result = mgr.get(filter=filter_xml)
        return True
    except Exception as e:
        logger.error(f"Request failed: {e}")
        return False


def run_sequential(args):
    """Run requests sequentially."""
    connect_kwargs = {
        "host": args.host,
        "port": args.port,
        "username": "user",
        "password": "pass",
        "hostkey_verify": False,
        "device_params": {"name": "default"},
        "timeout": 30,
    }

    # Add libssh flag if using libssh backend
    if args.backend == "libssh":
        connect_kwargs["use_libssh"] = True
    else:
        # Paramiko-specific options
        connect_kwargs["allow_agent"] = False
        connect_kwargs["look_for_keys"] = False

    with manager.connect(**connect_kwargs) as mgr:
        for i in range(args.count):
            run_get_request(mgr, args.size)


def main():
    parser = argparse.ArgumentParser(description="NETCONF ncclient benchmark")
    parser.add_argument("--host", default="localhost", help="NETCONF server host")
    parser.add_argument("--port", type=int, default=8830, help="NETCONF server port")
    parser.add_argument("--size", type=int, default=1024, help="Response size in bytes")
    parser.add_argument("--count", type=int, default=10, help="Number of requests")
    parser.add_argument(
        "--backend",
        default="paramiko",
        choices=["paramiko", "libssh"],
        help="SSH backend to use (default: paramiko)",
    )

    args = parser.parse_args()

    logger.info(f"Connecting to {args.host}:{args.port}")
    logger.info(f"Using backend: {args.backend}")
    logger.info(f"Running {args.count} requests with size={args.size}")

    start = time.time()

    try:
        run_sequential(args)
    except Exception as e:
        logger.error(f"Benchmark failed: {e}")
        sys.exit(1)

    duration = time.time() - start

    # Print timing to stdout for hyperfine
    print(f"{duration:.3f}")

    # Log stats to stderr
    throughput = args.count / duration
    data_transferred = (args.count * args.size) / (1024 * 1024)
    logger.info(f"Completed in {duration:.3f}s")
    logger.info(f"Throughput: {throughput:.2f} req/s")
    logger.info(f"Data transferred: {data_transferred:.2f} MB")


if __name__ == "__main__":
    main()
