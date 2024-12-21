#!/usr/bin/env bash

function yell () { >&2 echo "$*";  }
function die () { yell "$*"; exit 1; }
function try () { "$@" || die "Command failed: $*"; }

script_path="$( realpath "$0" )"
script_dir="$( dirname "$script_path" )"

function case1() {
    local input_path
    local output_path
    local distance
    local debug_svg

    input_path="$script_dir/pcb_edge_vertices.json"
    output_path="$script_dir/pcb_edge_vertices_expanded.json"
    distance="5.0"
    debug_svg="$script_dir/pcb_edge_vertices_expanded.svg"

    kbutil expand-vertices2 "$input_path" "$output_path" "$distance" --debug-svg="$debug_svg"
}

function case2() {
    local input_path
    local output_path
    local distance
    local debug_svg

    input_path="$script_dir/case2.json"
    output_path="$script_dir/case2_expanded.json"
    distance="5.0"
    debug_svg="$script_dir/case2_expanded.svg"

    kbutil expand-vertices2 "$input_path" "$output_path" "$distance" --debug-svg="$debug_svg"
    #kbutil expand-vertices "$input_path" "$output_path" "$distance" --debug-svg="$debug_svg"
}

case1
#case2
