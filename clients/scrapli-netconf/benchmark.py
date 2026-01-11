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

logging.getLogger("scrapli").setLevel(logging.WARNING)

logger = logging.getLogger(__name__)


def run_get_request(conn, size):
    """Execute a single get request (sync)."""


def run_sync(args):
    transport_options = {}

    # For system transport, disable PTY allocation for NETCONF
    if args.transport == "system":
        transport_options["netconf_force_pty"] = False
        transport_options["ssh_args"] = ["-T"]

    filter = f'<filter type="subtree"><size>{args.size}</size></filter>'

    start = time.time()
    conn = NetconfDriver(
        host=args.host,
        port=args.port,
        auth_username="user",
        auth_password="pass",
        auth_strict_key=False,
        ssh_config_file=False,
        transport=args.transport,
        transport_options=transport_options,
        timeout_socket=300,
        timeout_transport=300,
        timeout_ops=300,
    )

    conn.open()
    startup_time = time.time() - start

    start = time.time()
    for i in range(args.count):
        result = conn.get(filter_=filter)
    duration = time.time() - start

    conn.close()

    print_results(args, startup_time, duration)


async def run_async(args):
    filter = f'<filter type="subtree"><size>{args.size}</size></filter>'
    start = time.time()
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
    startup_time = time.time() - start

    start = time.time()
    for i in range(args.count):
        await conn.get(filter_=filter)
    duration = time.time() - start

    await conn.close()

    print_results(args, startup_time, duration)


def print_results(args, startup_time, duration):
    # impl,setup_ms,rpc_calls_ms,rpc
    print(
        "Python: scrapli-netconf ({}),{},{},{},{:.0f},{:.0f},{:.3f}".format(
            args.transport,
            args.framing,
            args.size,
            args.count,
            startup_time * 1000,
            duration * 1000,
            args.count / duration,
        )
    )


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
    parser.add_argument(
        "--framing",
        choices=["marked", "chunked"],
        help="Framing type",
        required=True,
    )

    args = parser.parse_args()

    logger.info(f"Connecting to {args.host}:{args.port}")
    logger.info(f"Using transport: {args.transport}")
    logger.info(f"Running {args.count} requests with size={args.size}")

    if args.transport == "asyncssh":
        asyncio.run(run_async(args))
    else:
        run_sync(args)


if __name__ == "__main__":
    main()
