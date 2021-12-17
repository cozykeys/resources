#!/usr/bin/env python3

import pdb
import argparse
import openpyscad as ops  # type: ignore

from lib.util import mkdir
from lib import TriadPiece, SwitchData


class Stats:
    def __init__(self, data):
        self.min = min(data)
        self.max = max(data)
        self.mid = self.min + ((self.max - self.min) / 2.0)

    def __str__(self):
        return "Min = {}, Max = {}, Mid = {}".format(
            round(self.min, 3), round(self.max, 3), round(self.mid, 3)
        )


def add_generate_case_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser("generate-case", help="Generate case using OpenSCAD")
    parser.set_defaults(func=generate_case)


def cube(x, y, z):
    (x, y, z) = (float(x), float(y), float(z))
    (tx, ty, tz) = (-(x / 2.0), -(y / 2.0), -(z / 2.0))
    return ops.Cube([x, y, z]).translate([tx, ty, tz])


def switch_cutout():
    c1 = cube(14.0, 14.0, 50.0)
    c2 = cube(15.6, 12.0, 50.0)
    return c1 + c2


def switch_cutouts(piece):
    switches = SwitchData(piece).get_switches()

    stats_x = Stats([s["x"] for s in switches])
    stats_y = Stats([s["y"] for s in switches])

    u = ops.Union()
    for s in switches:
        x = s["x"] - stats_x.mid
        y = -(s["y"] - stats_y.mid)
        u.append(switch_cutout().translate([x, y, 0.0]))
    return u


def switch_cutouts_union():
    u1 = (
        switch_cutouts(TriadPiece.LEFT)
        .translate([-100.0, 0.0, 0.0])
        .rotate([-10, 10, 0])
    )
    u2 = switch_cutouts(TriadPiece.CENTER)
    u3 = switch_cutouts(TriadPiece.RIGHT).translate([100.0, 0.0, 0.0]).rotate(10)

    return u1 + u2 + u3


def plate():
    c1 = cube(155, 155, 4.5).translate([-102.5, 13.0, 0.0]).rotate([0, 10, 0])
    c2 = cube(50, 155, 4.5).translate([0.0, 13.0, 0.0])
    c3 = cube(155, 155, 4.5).translate([102.5, 13.0, 0.0])
    return c1 + c2 + c3

    # return cube(375, 155, 4.5).translate([0.0, 13.0, 0.0])


def foo():

    # p = ops.Polygon(points=[
    # [ -100.0, -100.0 ],
    # [  100.0, -100.0 ],
    # [  100.0,  100.0 ],
    # [ -100.0,  100.0 ],
    # ])
    # x = p.linear_extrude(height=10)
    # x = x.translate([0.0, -5.0, 0.0])
    # c1 = cube(50, 50, 50)
    # return (x - c1)

    edge_vertices = [
        [29.775, 25.35],
        [86.925, 17.35],
        [105.975, 17.35],
        [146.451, 29.036],
        [163.125, 123.6],
        [163.125, 142.65],
        [125.025, 142.65],
        [67.875, 139.65],
        [29.775, 139.65],
    ]

    stats_x = Stats([v[0] for v in edge_vertices])
    stats_y = Stats([v[1] for v in edge_vertices])

    points = []
    for v in edge_vertices:
        x = v[0] - stats_x.mid
        y = -(v[1] - stats_y.mid)
        points.append([x, y])

    c1 = ops.Polygon(points=points).linear_extrude(4.5).translate([0.0, -2.25, 0.0])
    u1 = (
        switch_cutouts(TriadPiece.LEFT)
        .translate([0.0, 0.0, 0.0])
        .rotate([0.0, 0.0, 0.0])
    )
    return c1 - u1

    c1 = cube(150, 155, 4.5).translate([-75 - 30, 13.0, 0.0]).rotate([0.0, 0.0, 0.0])
    c2 = cube(60, 155, 4.5).translate([0.0, 13.0, 0.0]).rotate([0.0, 0.0, 0.0])
    c3 = cube(150, 155, 4.5).translate([75 + 30, 13.0, 0.0]).rotate([0.0, 0.0, 0.0])

    u1 = (
        switch_cutouts(TriadPiece.LEFT)
        .translate([-100.0, 0.0, 0.0])
        .rotate([0.0, 0.0, -10.0])
    )
    u2 = switch_cutouts(TriadPiece.CENTER)
    u3 = (
        switch_cutouts(TriadPiece.RIGHT)
        .translate([100.0, 0.0, 0.0])
        .rotate([0.0, 0.0, 10.0])
    )

    x1 = (c1 - u1).rotate([0.0, -10.0, 0.0])

    return x1 + (c2 - u2) + (c3 - u3)


def generate_case(args: argparse.Namespace) -> int:
    print(
        """
********************************************************************************
                    WARNING: THIS FILE IS CURRENTLY GARBAGE!
                     I was using it to learn how to script
                  OpenSCAD in Python and haven\'t cleaned it up.
********************************************************************************
"""
    )

    # p = plate()

    # switch_cutouts = switch_cutouts_union()
    # mkdir('./temp')
    # (p - switch_cutouts).write('./temp/triad.scad')

    foo().write("./temp/triad.scad")

    return 0
