#!/usr/bin/env bash

COLOR_RED='\033[0;31m'
COLOR_NONE='\033[0m'

try()
{
    "$@"
    local ret_val=$?
  
    if [ ! $ret_val -eq 0 ]; then
        echo -e "${COLOR_RED}Command failed ($ret_val):${COLOR_NONE} $*"
        exit 1
    fi
}

argus_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
kbutil_dir="$( dirname "$argus_dir" )/kbutil"

kbutil_dll="$kbutil_dir/build/KbUtil.Console/bin/Release/kbutil.dll"

input="$argus_dir/argus.xml"
output="$argus_dir/case"

svg_opener="inkscape"

options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

try dotnet "$kbutil_dll" gen-svg $options "$input" "$output"
#dotnet "$kbutil_dll" gen-key-bearings "$input" "./keys.json" --debug-svg="./temp.svg"

try "$svg_opener" "$output/argus_Switch.svg" 

