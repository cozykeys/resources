#!/usr/bin/env bash

# TODO: This script can be deleted once adequate feature parity has been
# accomplished in SVG generation. For now it's here to make comparing the
# results of kbutil and kb easier.

function yell () { >&2 echo "$*";  }
function die () { yell "$*"; exit 1; }
function try () { "$@" || die "Command failed: $*"; }

script_path="$( realpath "$0" )"
script_dir="$( dirname "$script_path" )"

kbutil gen-svg \
    "$(pwd)/pkg/unmarshal/test_data/bloomer_v4.xml" \
    "$(pwd)/temp2" \
    --visual-switch-cutouts \
    --keycap-overlays \
    --keycap-legends
    #--squash

./bin/kb svg \
    "$(pwd)/pkg/unmarshal/test_data/bloomer_v4.xml" \
    "$(pwd)/temp"
