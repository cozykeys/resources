#!/usr/bin/env python

import argparse

from commands import add_command_parsers
from lib.log import Log


def parse_args():
    parser = argparse.ArgumentParser(description="Python Keyboard Tools CLI")
    parser.add_argument(
        "-l",
        "--log-level",
        default="info",
        help="Logging level to run with (debug, info, warn, error, crit)",
    )

    add_command_parsers(parser)

    # TODO: Bash auto-completion

    args = parser.parse_args()
    if "func" not in args:
        parser.print_help()
        exit(1)
    return args


def main():
    args = parse_args()
    Log.init(args.log_level)
    args.func(args)


if __name__ == "__main__":
    main()
