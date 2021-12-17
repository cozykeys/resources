#!/usr/bin/env python3

import os
import argparse
from lib.util import sh, mkdir, get_triad_repo_dir
from lib.kbutil import get_kbutil_dll_path


def add_ponoko_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser("ponoko")
    parser.add_argument("--open", action="store_true")
    parser.set_defaults(func=ponoko)


def generate_svg(kbutil_dll_path, input_path, output_path):
    return sh(["dotnet", kbutil_dll_path, "gen-svg", input_path, output_path])


def ponoko(args: argparse.Namespace) -> int:
    kbutil_dll = get_kbutil_dll_path()
    output_path = os.path.join(get_triad_repo_dir(), "case", "ponoko")
    mkdir(output_path)

    for piece in ["left", "center", "right"]:
        print(piece)

        input_path = "layout/triad_{}.xml".format(piece)
        ret = generate_svg(kbutil_dll, input_path, output_path)
        if ret != 0:
            print("Failed to generate SVG")
            return ret

    if args.open:
        sh(
            [
                "inkscape",
                os.path.join(output_path, "triad_left.svg"),
                os.path.join(output_path, "triad_right.svg"),
                os.path.join(output_path, "triad_center.svg"),
            ]
        )

    return 0
