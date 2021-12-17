#!/usr/bin/env python3

import pdb
import argparse
import openpyscad as ops  # type: ignore

from lib.util import mkdir
from lib import TriadPiece, SwitchData

# Useful Data
#
# Left
#   Left edge of left plate:
#     (-69.675, 59.65) -> (-69.675, -54.259)
# 
#   Tenting Stand Holes:
#     (48.825, 34.875)
#     (48.825, 130.125)
#     (144.075, 34.875)
#     (144.075, 130.125)
# 
#   Tenting Stand Holes (Centered):
#     (-47.625, -45.125)
#     (-47.625, 50.125)
#     (47.625, -45.125)
#     (47.625, 50.125)


def add_generate_stand_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser("generate-stand")
    parser.set_defaults(func=generate_stand)


def cube(x, y, z):
    (x, y, z) = (float(x), float(y), float(z))
    (tx, ty, tz) = (-(x / 2.0), -(y / 2.0), -(z / 2.0))
    return ops.Cube([x, y, z]).translate([tx, ty, tz])


def generate_risers_side():
    # These should go from -69.675, the left edge of the case, to 51.625, 4mm
    # past the nubs
    risers_union = ops.Union()
    risers = [
        cube(121.3, 10, 25).translate([-9.025, -45.125, 0.5]),
        cube(121.3, 10, 25).translate([-9.025, 50.125, 0.5]),
    ]
    for r in risers:
        risers_union.append(r)
    return risers_union


def generate_nubs_side():
    nubs_union = ops.Union()
    nubs = [
        # 31.0 mm makes the nub stick out 3mm from the riser
        ops.Cylinder(h=31.0, r=1, center=True).translate([-47.625, -45.125, 0.0]),
        ops.Cylinder(h=31.0, r=1, center=True).translate([-47.625, 50.125, 0.0]),
        ops.Cylinder(h=31.0, r=1, center=True).translate([47.625, -45.125, 0.0]),
        ops.Cylinder(h=31.0, r=1, center=True).translate([47.625, 50.125, 0.0]),
    ]
    for n in nubs:
        nubs_union.append(n)
    return nubs_union


def generate_braces_side():
    # These should range from Y = -45.125 to Y = 50.125
    braces_union = ops.Union()
    braces = [
        cube(10, 95.25, 10).translate([0.0, 2.5, 5.0]),
        cube(10, 95.25, 10).translate([40.0, 2.5, 5.0]),
    ]
    for b in braces:
        braces_union.append(b)
    return braces_union

def generate_side():
    # Generate the risers that the keyboard will sit on
    risers = generate_risers_side()

    # Add nubs that go into the case holes to keep it in place
    nubs = generate_nubs_side()

    # Union and rotate to a 10* grade
    supports = (risers + nubs).rotate([0.0, -10.0, 0.0])

    # Add brace supports that join the two risers
    braces = generate_braces_side()

    return supports + braces


def generate_risers_center():
    # These should go from Y = -62.125 to Y = 63.125

    # TODO: Figure out correct Z value
    risers_union = ops.Union()
    risers = [
        cube(10, 125.25, 40).translate([-25.05, 0.5, 0.0]),
        cube(10, 125.25, 40).translate([25.05, 0.5, 0.0]),
    ]
    for r in risers:
        risers_union.append(r)
    return risers_union


def generate_nubs_center():
    # Tenting holes:
    #    (-25.05, -42.125),
    #    (25.05, -42.125),
    #    (-25.05, 53.125),
    #    (25.05, 53.125),

    nubs_union = ops.Union()
    nubs = [
        # 31.0 mm makes the nub stick out 3mm from the riser
        ops.Cylinder(h=46.0, r=1, center=True).translate([-25.05, -45.125, 0.0]),
        ops.Cylinder(h=46.0, r=1, center=True).translate([-25.05, 53.125, 0.0]),
        ops.Cylinder(h=46.0, r=1, center=True).translate([25.05, -45.125, 0.0]),
        ops.Cylinder(h=46.0, r=1, center=True).translate([25.05, 53.125, 0.0]),
    ]
    for n in nubs:
        nubs_union.append(n)
    return nubs_union


def generate_braces_center():
    braces_union = ops.Union()
    braces = [
        cube(50.1, 10, 10).translate([0.0, -30.0, 5.0]),
        cube(50.1, 10, 10).translate([0.0, 30.0, 5.0]),
    ]
    for b in braces:
        braces_union.append(b)
    return braces_union


def generate_center():
    # Generate the risers that the keyboard will sit on
    risers = generate_risers_center()

    # Add nubs that go into the case holes to keep it in place
    nubs = generate_nubs_center()

    # Union
    supports = (risers + nubs)

    # Add brace supports that join the two risers
    braces = generate_braces_center()

    return supports + braces


def generate_stand(args: argparse.Namespace) -> int:
    base = cube(500, 500, 200).translate([0.0, 0.0, -100.0])

    left = generate_side()
    (left - base).write("./temp/triad_stand_left.scad")

    center = generate_center()
    (center - base).write("./temp/triad_stand_center.scad")

    right = left.mirror([1.0, 0.0, 0.0])
    (right - base).write("./temp/triad_stand_right.scad")

    left = left.translate([-120.0, 0.0, 0.0]).rotate([0.0, 0.0, 10.0])
    right = right.translate([120.0, 0.0, 0.0]).rotate([0.0, 0.0, -10.0])
    center = center.translate([0.0, -10.0, 0.0])
    ((left + center + right) - base).write("./temp/triad_stand_render.scad")

    return 0
