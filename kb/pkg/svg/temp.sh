#!/usr/bin/env bash

function yell () { >&2 echo "$*";  }
function die () { yell "$*"; exit 1; }
function try () { "$@" || die "Command failed: $*"; }

script_path="$( realpath "$0" )"
script_dir="$( dirname "$script_path" )"

touch layer_writer.go

touch circle_writer.go
touch group_writer.go
touch key_writer.go
touch legend_writer.go
touch spacer_writer.go
touch stack_writer.go
touch text_writer.go
touch path_writer.go
