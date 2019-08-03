#!/usr/bin/env bash

repo_root_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
kbutil_dll="$repo_root_dir/build/KbUtil.Console/bin/Release/kbutil.dll"
goobie_dir="$repo_root_dir/../goobie"
input="$goobie_dir/goobie.xml"
output="$goobie_dir/case"

open_with="inkscape"

options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

dotnet "$kbutil_dll" gen-svg $options "$input" "$output"

"$open_with" "$output/goobie_Switch.svg" 

