#!/usr/bin/env bash


radix_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
kbutil_dir="$( dirname "$radix_dir" )/kbutil"

kbutil_dll="$kbutil_dir/build/KbUtil.Console/bin/Release/kbutil.dll"

input="$radix_dir/radix_next.xml"
output="$radix_dir/case"

svg_opener="inkscape"

options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

dotnet "$kbutil_dll" gen-svg $options "$input" "$output"
#dotnet "$kbutil_dll" gen-key-bearings "$input" "./keys.json" --debug-svg="./temp.svg"

"$svg_opener" "$output/radix_Switch.svg" 

