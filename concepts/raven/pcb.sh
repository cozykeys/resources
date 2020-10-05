#!/usr/bin/env bash

raven_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
kbutil_dir="$( dirname "$raven_dir" )/kbutil"

kbutil_dll="$kbutil_dir/build/KbUtil.Console/bin/Release/kbutil.dll"

pcb_dir="$raven_dir/pcb"

input="$raven_dir/switches.json"
output="$pcb_dir/raven.kicad_pcb"

(
    cd "$kbutil_dir"
    make
)

mkdir -p "$pcb_dir"
dotnet "$kbutil_dll" gen-pcb "raven" "$input" "$output"

