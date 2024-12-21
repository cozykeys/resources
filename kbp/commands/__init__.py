#!/usr/bin/env python

import argparse


from .scratch import add_scratch_parser


def add_command_parsers(parser: argparse.ArgumentParser) -> None:
    subparsers = parser.add_subparsers(help="commands")

    add_scratch_parser(subparsers)
