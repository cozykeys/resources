#!/usr/bin/env bash

kbutil_exe="src/KbUtil.Console/bin/Release/netcoreapp2.0/linux-x64/kbutil"
input="$HOME/src/cozykeys/bloomer/bloomer.xml"
output="$HOME/src/cozykeys/svgs4"

open_with="inkscape"
#open_with="chromium-browser"

#options="--visual-switch-cutouts"
options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

"$kbutil_exe" gen-svg $options "$input" "$output"

"$open_with" "$output/bloomer_Switch.svg"

