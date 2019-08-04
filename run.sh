#!/usr/bin/env bash

repo_root_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
kbutil_dll="$repo_root_dir/build/KbUtil.Console/bin/Release/kbutil.dll"
radix_dir="$repo_root_dir/../radix"
input="$radix_dir/radix.xml"
output="$radix_dir/case"

open_with="inkscape"

options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

dotnet "$kbutil_dll" gen-svg $options "$input" "$output"
dotnet "$kbutil_dll" gen-key-bearings "$input" "./keys.json" --debug-svg="./temp.svg"

#"$open_with" "$output/radix_Switch.svg" 

