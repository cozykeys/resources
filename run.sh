#!/usr/bin/env bash

kbutil_exe="src/KbUtil.Console/bin/Release/netcoreapp2.0/linux-x64/kbutil"
input="$HOME/src/cozykeys/bloomer/bloomer.xml"
#input="$HOME/src/cozykeys/bloomer/layouts/rev3/default/base.xml"
#input="$HOME/src/cozykeys/bloomer/layouts/rev3/default/fn.xml"
#input="$HOME/src/cozykeys/bloomer/layouts/rev3/default/num.xml"
output="$HOME/src/cozykeys/svgs4"

open_with="inkscape"
#open_with="chromium-browser"

options="--visual-switch-cutouts"
#options="--visual-switch-cutouts --keycap-overlays --keycap-legends"

"$kbutil_exe" gen-svg $options "$input" "$output"

"$open_with" "$output/bloomer_Switch.svg" &

#"$open_with" "$output/bloomer_Num.svg" &
#"$open_with" "$output/bloomer_Fn.svg" &
#"$open_with" "$output/bloomer_Base.svg" &

wait
