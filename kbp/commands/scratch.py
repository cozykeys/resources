#!/usr/bin/env python

import argparse
import os
import math

from lib.log import Log
from lib.geometry_2d import Segment2D, Vector2D
from lib.error import NotYetImplementedException

class Angle2D:
    @staticmethod
    def from_degrees(degrees):
        return Angle2D(math.radians(degrees))

    @staticmethod
    def from_radians(radians):
        return Angle2D(radians)

    @staticmethod
    def from_vector():
        raise NotYetImplementedException()

    def __init__(self, radians):
        # TODO: Modulo this or whatever so it's in the range -math.PI to math.PI
        self._radians = radians

    def to_vector(self):
        return Vector2D(
            math.cos(self._radians),
            math.sin(self._radians),
        )



def add_scratch_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser(
        "scratch",
        help="Scratch pad for WIP functionality",
    )
    parser.set_defaults(func=cmd_scratch)


def project_from_point(
    point: Vector2D, direction: Vector2D, distance: float
) -> Vector2D:
    diff = Vector2D(direction.x * distance, direction.y * distance)
    return Vector2D(point.x + diff.x, point.y + diff.y)


def cmd_scratch(args: argparse.Namespace) -> None:
    # First, get the point 5mm away that we want the segment to intersect with
    point = Vector2D(170.900, 153.626)
    direction = Angle2D.from_degrees(190).to_vector()
    distance = 5.0
    projection = project_from_point(point, direction, distance)
    Log.info(f"Projected point: {projection}")

    # Next, calculate what the value of X is for the segment at the same Y as
    # the projected point.
    segment = Segment2D(Vector2D(160.209, 173.884), Vector2D(180.239, 60.287))
    #segment = Segment2D(Vector2D(162.209, 151.143), Vector2D(182.239, 37.546))
    line = segment.to_line()
    x = line.get_x_at_y(projection.y)
    Log.info(f"Point on segment: {Vector2D(x, projection.y)}")

    # Get how much the segment needs to move to intersection the projected
    # point
    dx = projection.x - x
    Log.info(f"Segment needs to be moved by {dx} on X axis")
    

