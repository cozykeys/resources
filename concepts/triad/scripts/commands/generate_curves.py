#!/usr/bin/env python3

import json
import argparse
from typing import List, Dict, Tuple
from lib import SvgStyle, SvgWriter, Polygon2D, Vector2D, Segment2D, Curve2D


SVG_STYLE_VEC: Dict[str, str] = {
    "fill-opacity": "1.0",
    "fill": "#7f007f",  # Purple
    "stroke": "none",
}
SVG_STYLE_POLY: Dict[str, str] = {
    "fill-opacity": "0.0",
    "stroke": "#7f007f",  # Purple
    "stroke-width": "0.25",
}
SVG_STYLE_CURVE: Dict[str, str] = {
    "fill-opacity": "0.0",
    "stroke": "#007f7f",  # Teal
    "stroke-width": "0.25",
}


def add_generate_curves_parser(subparsers: argparse._SubParsersAction) -> None:
    parser = subparsers.add_parser(
        "generate-curves", help="Given a path (Set of vertices), generate smooth curves"
    )
    parser.add_argument("input_path")
    parser.add_argument("output_path")
    parser.add_argument("distance", type=float)
    parser.add_argument("--debug-svg", type=str, default="")
    parser.set_defaults(func=generate_curves)


def project_from_point(
    point: Vector2D, direction: Vector2D, distance: float
) -> Vector2D:
    diff = Vector2D(direction.x * distance, direction.y * distance)
    return Vector2D(point.x + diff.x, point.y + diff.y)


def poly_curves(polygon: Polygon2D, distance: float) -> List[Curve2D]:
    curves: List[Curve2D] = []
    for curr in range(0, len(polygon.vertices)):
        prev = curr - 1 if curr > 0 else len(polygon.vertices) - 1
        next = curr + 1 if curr < len(polygon.vertices) - 1 else 0

        s1 = Segment2D(polygon.vertices[curr], polygon.vertices[prev])
        start = project_from_point(polygon.vertices[curr], s1.direction(), distance)

        s2 = Segment2D(polygon.vertices[curr], polygon.vertices[next])
        end = project_from_point(polygon.vertices[curr], s2.direction(), distance)

        curves.append(Curve2D(start, end, polygon.vertices[curr]))

    return curves


def print_layout_path_components(curves: List[Curve2D]) -> None:
    lines = [
        "<AbsoluteMoveTo>",
        '    <EndPoint X="{}" Y="{}" />'.format(
            round(curves[0].start.x, 3), round(curves[0].start.y, 3)
        ),
        "</AbsoluteMoveTo>",
    ]

    for i in range(0, len(curves)):
        curr = curves[i]
        next = curves[(i + 1) % len(curves)]
        lines += [
            "<AbsoluteQuadraticCurveTo>",
            '    <EndPoint X="{}" Y="{}" />'.format(
                round(curr.end.x, 3), round(curr.end.y, 3)
            ),
            '    <ControlPoint X="{}" Y="{}" />'.format(
                round(curr.control.x, 3), round(curr.control.y, 3)
            ),
            "</AbsoluteQuadraticCurveTo>",
            "<AbsoluteLineTo>",
            '    <EndPoint X="{}" Y="{}" />'.format(
                round(next.start.x, 3), round(next.start.y, 3)
            ),
            "</AbsoluteLineTo>",
        ]

    print("\n".join(lines))


def generate_curves(args: argparse.Namespace) -> int:
    json_data: List[Tuple[float, float]] = []
    with open(args.input_path, "r") as f:
        json_data = json.loads(f.read())

    poly = Polygon2D.from_json(json_data)
    curves = poly_curves(poly, args.distance)

    svg_writer = SvgWriter()
    svg_writer.append_element(poly, SvgStyle(SVG_STYLE_POLY))
    for v in poly.vertices:
        svg_writer.append_element(v, SvgStyle(SVG_STYLE_VEC))
    for c in curves:
        svg_writer.append_element(c, SvgStyle(SVG_STYLE_CURVE))
    if args.debug_svg != "":
        print('Writing debug svg to "{}"'.format(args.debug_svg))
        svg_writer.write_to_file(args.debug_svg)
    else:
        svg_writer.write_to_stdout()

    print_layout_path_components(curves)

    with open(args.output_path, "w") as f:
        f.write(json.dumps([c.to_json() for c in curves], indent=4))

    return 0
