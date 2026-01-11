#!/usr/bin/env python3

import argparse
import logging
import sys
import time

from ncclient import manager

# Configure app logging
logging.basicConfig(
    level=logging.INFO, format="%(levelname)s:%(message)s", stream=sys.stderr
)

# Silence ncclient's verbose logging
logging.getLogger("ncclient").setLevel(logging.WARNING)

logger = logging.getLogger(__name__)


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
    parser.add_argument(
        "--framing",
        choices=["marked", "chunked"],
        help="Framing type",
        required=True,
    )

    args = parser.parse_args()

    logger.info(f"Connecting to {args.host}:{args.port}")
    logger.info(f"Using backend: {args.backend}")
    logger.info(f"Running {args.count} requests with size={args.size}")

    connect_kwargs = {}
    if args.backend == "libssh":
        connect_kwargs["use_libssh"] = True
    else:
        # Paramiko-specific options
        connect_kwargs["allow_agent"] = False
        connect_kwargs["look_for_keys"] = False

    start = time.time()
    mgr = manager.connect(
        host=args.host,
        port=args.port,
        username="admin",
        password="admin",
        hostkey_verify=False,
        **connect_kwargs,
    )
    startup_time = time.time() - start

    filter = f'<filter type="subtree"><size>{args.size}</size></filter>'

    start = time.time()
    for i in range(args.count):
        mgr.get(filter=filter)
    duration = time.time() - start
    mgr.close_session()

    # impl,setup_ms,rpc_calls_ms,rpc
    print(
        "Python: ncclient ({}),{},{},{},{:.0f},{:.0f},{:.3f}".format(
            args.backend,
            args.framing,
            args.size,
            args.count,
            startup_time * 1000,
            duration * 1000,
            args.count / duration,
        )
    )


if __name__ == "__main__":
    main()
