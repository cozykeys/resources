#!/usr/bin/env python3

import os
import argparse
from lib.util import sh, mkdir
from lib.kbutil import get_kbutil_dll_path


def add_render_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser(
        "render", help="Renders the triad keyboard onto an SVG"
    )
    parser.add_argument("--open", action="store_true")
    parser.set_defaults(func=render)


def generate_svg(kbutil_dll_path, input_path, output_path):
    return sh(
        [
            "dotnet",
            kbutil_dll_path,
            "gen-svg",
            "--visual-switch-cutouts",
            "--keycap-overlays",
            "--keycap-legends",
            # "--squash",
            input_path,
            output_path,
        ]
    )


def generate_bearings(kbutil_dll_path, input_path, output_path, debug_path):
    return sh(
        [
            "dotnet",
            kbutil_dll_path,
            "gen-key-bearings",
            input_path,
            output_path,
            "--debug-svg={}".format(debug_path),
        ]
    )


def render(args: argparse.Namespace) -> int:
    kbutil_dll = get_kbutil_dll_path()
    output_path = "case"

    for piece in ["left", "center", "right"]:
        input_path = "layout/triad_{}.xml".format(piece)
        ret = generate_svg(kbutil_dll, input_path, output_path)
        if ret != 0:
            print("Failed to generate SVG")
            return ret

        bearings_path = "./temp/bearings_{}.json".format(piece)
        bearings_debug_path = "./temp/bearings_{}.svg".format(piece)

        mkdir(os.path.dirname(bearings_path))
        ret = generate_bearings(
            kbutil_dll, input_path, bearings_path, bearings_debug_path
        )
        if ret != 0:
            print("Failed to switch bearings")
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
