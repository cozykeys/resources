#!/usr/bin/env python3

import json
import argparse
import math
from typing import List, Dict, Tuple
from lib import SvgStyle, SvgWriter, Vector2D, Segment2D, Polygon2D, get_triad_repo_dir


SVG_STYLE_POLY: Dict[str, str] = {
    "fill-opacity": "0.0",
    "stroke": "#7f007f",  # Purple
    "stroke-width": "0.25",
}
SVG_STYLE_VEC: Dict[str, str] = {
    "fill-opacity": "1.0",
    "fill": "#7f007f",  # Purple
    "stroke": "none",
}
SVG_STYLE_VEC_1: Dict[str, str] = {
    "fill-opacity": "1.0",
    "fill": "#ff0000",  # Purple
    "stroke": "none",
}
SVG_STYLE_HOLE: Dict[str, str] = {
    "fill-opacity": "1.0",
    "fill": "#007f7f",  # Teal
    "stroke": "none",
}


def add_scratch_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser(
        "scratch", help="Scratchpad, dumping ground for temporary code"
    )
    parser.set_defaults(func=scratch)


def scratch(args: argparse.Namespace) -> int:
    triad_dir = get_triad_repo_dir()

    for piece in ["left", "right"]:
        print(piece)

        if piece == "left":
            conn_pos = Vector2D(155.228, 110.485)
        if piece == "right":
            conn_pos = Vector2D(37.672, 110.485)

        conn_poly = Polygon2D(
            [
                Vector2D(-8.0, -0.75),
                Vector2D(8.0, -0.75),
                Vector2D(8.0, 5.25),
                Vector2D(-8.0, 5.25),
            ]
        )

        if piece == "left":
            degrees = -100
        if piece == "right":
            degrees = 100

        poly = conn_poly.rotated_around(math.radians(degrees), Vector2D(0.0, 0.0))
        poly = poly.translated(conn_pos)

        poly_expanded = poly.expanded(-1.0)

        s1 = Segment2D(poly_expanded.vertices[1], poly_expanded.vertices[2])
        s2 = Segment2D(poly_expanded.vertices[0], poly_expanded.vertices[3])
        if piece == "left":
            s3 = Segment2D(
                Vector2D(52.631 + 96.45, -53.327 + 80),
                Vector2D(69.154 + 96.45, 40.384 + 80),
            )
        elif piece == "right":
            s3 = Segment2D(
                Vector2D(-52.631 + 96.45, -53.327 + 80),
                Vector2D(-69.154 + 96.45, 40.384 + 80),
            )

        v1 = s1.intersection(s3)
        v2 = s2.intersection(s3)

        (dx, dy) = (96.45, 80)
        points = [
            Vector2D(v1.x - dx, v1.y - dy),
            Vector2D(s1.start.x - dx, s1.start.y - dy),
            Vector2D(s2.start.x - dx, s2.start.y - dy),
            Vector2D(v2.x - dx, v2.y - dy),
        ]

        print("Polygon:")
        for p in poly.vertices:
            print("    {}".format(p))

        svg_writer = SvgWriter()
        svg_writer.append_element(poly, SvgStyle(SVG_STYLE_POLY))
        for v in poly.vertices:
            svg_writer.append_element(v, SvgStyle(SVG_STYLE_VEC))
        svg_writer.append_element(conn_pos, SvgStyle(SVG_STYLE_VEC))
        svg_writer.append_element(poly_expanded, SvgStyle(SVG_STYLE_POLY))
        for v in poly_expanded.vertices:
            svg_writer.append_element(v, SvgStyle(SVG_STYLE_VEC))
        svg_writer.append_element(s3, SvgStyle(SVG_STYLE_POLY))
        svg_writer.append_element(v1, SvgStyle(SVG_STYLE_VEC))
        svg_writer.append_element(v2, SvgStyle(SVG_STYLE_VEC))
        for p in points:
            svg_writer.append_element(p, SvgStyle(SVG_STYLE_VEC_1))
        svg_writer.write_to_file("{}/temp/ffc_{}_debug.svg".format(triad_dir, piece))

        print("Points:")
        for p in points:
            print("    {}".format(p))

    # degrees =
    # radians = math.radians(degrees)
    # Vector2D(math.cos(radians), math.sin(radians)).normalized()

    return 0
