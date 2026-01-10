#!/usr/bin/env python3

import argparse
import asyncio
import logging
import sys
import time

from scrapli_netconf import AsyncNetconfDriver, NetconfDriver

# Configure app logging
logging.basicConfig(
    level=logging.INFO, format="%(levelname)s:%(message)s", stream=sys.stderr
)

# Silence scrapli's verbose logging
logging.getLogger("scrapli").setLevel(logging.WARNING)

logger = logging.getLogger(__name__)


def create_filter(size):
    """Create a NETCONF filter with size attribute."""
    return f'<filter type="subtree"><size>{size}</size></filter>'


def run_get_request(conn, size):
    """Execute a single get request (sync)."""
    filter_xml = create_filter(size)
    try:
        result = conn.get(filter_=filter_xml)
        return True
    except Exception as e:
        logger.error(f"Request failed: {e}")
        return False


async def run_get_request_async(conn, size):
    """Execute a single get request (async)."""
    filter_xml = create_filter(size)
    try:
        result = await conn.get(filter_=filter_xml)
        return True
    except Exception as e:
        logger.error(f"Request failed: {e}")
        return False


def run_sequential_sync(args):
    """Run requests sequentially (sync transports)."""
    transport_options = {}

    # For system transport, disable PTY allocation for NETCONF
    if args.transport == "system":
        transport_options["netconf_force_pty"] = False

    conn = NetconfDriver(
        host=args.host,
        port=args.port,
        auth_username="user",
        auth_password="pass",
        auth_strict_key=False,
        ssh_config_file=False,
        transport=args.transport,
        transport_options=transport_options,
        timeout_socket=30,
        timeout_transport=30,
        timeout_ops=30,
    )

    conn.open()
    try:
        for i in range(args.count):
            run_get_request(conn, args.size)
    finally:
        conn.close()


async def run_sequential_async(args):
    """Run requests sequentially (async transport)."""
    conn = AsyncNetconfDriver(
        host=args.host,
        port=args.port,
        auth_username="user",
        auth_password="pass",
        auth_strict_key=False,
        ssh_config_file=False,
        transport=args.transport,
        timeout_socket=30,
        timeout_transport=30,
        timeout_ops=30,
    )

    await conn.open()
    try:
        for i in range(args.count):
            await run_get_request_async(conn, args.size)
    finally:
        await conn.close()


def main():
    parser = argparse.ArgumentParser(description="NETCONF scrapli benchmark")
    parser.add_argument("--host", default="localhost", help="NETCONF server host")
    parser.add_argument("--port", type=int, default=8830, help="NETCONF server port")
    parser.add_argument("--size", type=int, default=1024, help="Response size in bytes")
    parser.add_argument("--count", type=int, default=10, help="Number of requests")
    parser.add_argument(
        "--transport",
        default="paramiko",
        choices=["system", "paramiko", "ssh2", "asyncssh"],
        help="SSH transport to use (default: paramiko)",
    )

    args = parser.parse_args()

    logger.info(f"Connecting to {args.host}:{args.port}")
    logger.info(f"Using transport: {args.transport}")
    logger.info(f"Running {args.count} requests with size={args.size}")

    start = time.time()

    try:
        if args.transport == "asyncssh":
            asyncio.run(run_sequential_async(args))
        else:
            run_sequential_sync(args)
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
