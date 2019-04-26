#!/usr/bin/env bash

kbutil_exe="dotnet $HOME/src/cozykeys/kbutil/src/KbUtil.Console/bin/Debug/netcoreapp2.0/kbutil.dll"
input="$HOME/src/cozykeys/bloomer_rev4_switches3.json"
output="$HOME/src/cozykeys/bloomer_rev4.kicad_pcb/bloomer"

open_with="kicad"

dotnet "/home/pewing/src/cozykeys/kbutil/src/KbUtil.Console/bin/Debug/netcoreapp2.0/kbutil.dll" gen-pcb "$input" "$output"

#"$open_with" "$output/bloomer_Switch.svg"

