#!/usr/bin/env bash


raven_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
kbutil_dir="$( dirname "$raven_dir" )/kbutil"

kbutil_dll="$kbutil_dir/build/KbUtil.Console/bin/Release/kbutil.dll"

input="$raven_dir/raven.xml"
output="$raven_dir/case"

svg_opener="inkscape"

options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

dotnet "$kbutil_dll" gen-svg $options "$input" "$output"
dotnet "$kbutil_dll" gen-key-bearings "$input" "./keys.json" --debug-svg="./temp.svg"

"$svg_opener" "$output/raven_switch.svg" 

