import argparse


def add_command_parsers(subparsers: argparse._SubParsersAction) -> None:
    from .center_polygon import add_center_polygon_parser

    add_center_polygon_parser(subparsers)

    from .expand_vertices import add_expand_vertices_parser

    add_expand_vertices_parser(subparsers)

    from .generate_case import add_generate_case_parser

    add_generate_case_parser(subparsers)

    from .generate_curves import add_generate_curves_parser

    add_generate_curves_parser(subparsers)

    from .generate_stand import add_generate_stand_parser

    add_generate_stand_parser(subparsers)

    from .generate_tenting_holes import add_generate_tenting_holes_parser

    add_generate_tenting_holes_parser(subparsers)

    from .ponoko import add_ponoko_parser

    add_ponoko_parser(subparsers)

    from .render import add_render_parser

    add_render_parser(subparsers)

    from .scratch import add_scratch_parser

    add_scratch_parser(subparsers)
