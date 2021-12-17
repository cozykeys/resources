#!/usr/bin/env bash

files="fp-info-cache
sym-lib-table
triad_right-cache.lib
triad_right.kicad_pcb
triad_right.kicad_pcb-bak
triad_right.net
triad_right.pro
triad_right-rescue.dcm
triad_right-rescue.lib
triad_right.sch
triad_right.sch-bak"

for f in $files; do
    sed -e 's/triad_left/triad_right/g' "$f" > "$f.new"
    mv "$f.new" "$f"
done

