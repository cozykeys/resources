#!/usr/bin/env bash

SCRIPTPATH="$( realpath $0 )"
SCRIPTDIR="$( dirname "$SCRIPTPATH" )"

triad_dir="$SCRIPTDIR"

"$triad_dir/scripts/triad.py" \
    expand-vertices \
    "$triad_dir/temp/pcb_edge_vertices_left.json" \
    "$triad_dir/temp/pcb_edge_vertices_left_expanded.json" \
    -5.0 \
    --debug-svg="$triad_dir/temp/debug.svg"
