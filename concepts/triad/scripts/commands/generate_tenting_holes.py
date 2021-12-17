#!/usr/bin/env python3

import json
import argparse
import math
from typing import List, Dict, Tuple
from lib import get_triad_repo_dir, SwitchData, TriadPiece


def add_generate_tenting_holes_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser("generate-tenting-holes")
    parser.set_defaults(func=generate_tenting_holes)


def generate_tenting_holes(args: argparse.Namespace) -> int:
    # triad_dir = get_triad_repo_dir()

    for piece in [TriadPiece.LEFT, TriadPiece.CENTER, TriadPiece.RIGHT]:
        print(piece)
        sd = SwitchData(piece)

        dx = (19.05 * 5.0) + 9.525
        holes = []
        if piece == TriadPiece.LEFT:
            holes += [sd.get_midpoint((0, 0), (0, 1))]
            holes += [sd.get_midpoint((5, 0), (5, 1))]
            holes += [
                (round(sd.get_switch_by_coord(0, 5)["x"] + 9.525, 3), holes[0][1])
            ]
            holes += [
                (round(sd.get_switch_by_coord(5, 5)["x"] + 9.525, 3), holes[1][1])
            ]
        elif piece == TriadPiece.CENTER:
            holes += [
                (
                    sd.get_switch_by_coord(0, 6)["x"] - 6,
                    sd.get_switch_by_coord(0, 6)["y"],
                )
            ]
            holes += [
                (
                    sd.get_switch_by_coord(0, 8)["x"] + 6,
                    sd.get_switch_by_coord(0, 8)["y"],
                )
            ]
            holes += [
                (
                    sd.get_switch_by_coord(5, 6)["x"] - 6,
                    sd.get_switch_by_coord(5, 6)["y"],
                )
            ]
            holes += [
                (
                    sd.get_switch_by_coord(5, 8)["x"] + 6,
                    sd.get_switch_by_coord(5, 8)["y"],
                )
            ]
            print("RESET: {}".format(sd.get_midpoint((0, 7), (1, 7))))
        elif piece == TriadPiece.RIGHT:
            holes += [sd.get_midpoint((0, 14), (0, 13))]
            holes += [sd.get_midpoint((5, 14), (5, 13))]
            holes += [
                (round(sd.get_switch_by_coord(0, 9)["x"] - 9.525, 3), holes[0][1])
            ]
            holes += [
                (round(sd.get_switch_by_coord(5, 9)["x"] - 9.525, 3), holes[1][1])
            ]
        else:
            raise Exception("Invalid piece")

        for h in holes:
            print("    {}".format(h))

        # TODO

    return 0
