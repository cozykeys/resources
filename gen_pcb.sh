#!/usr/bin/env bash

kbutil_exe="dotnet $HOME/src/cozykeys/kbutil/src/KbUtil.Console/bin/Debug/netcoreapp2.0/kbutil.dll"
input="$HOME/src/cozykeys/radix/keys.json"
output="$HOME/src/cozykeys/radix/pcb/test.kicad_pcb"

open_with="kicad"

dotnet "./build/KbUtil.Console/bin/Release/kbutil.dll" gen-pcb "$input" "$output"

#"$open_with" "$output/bloomer_Switch.svg"

