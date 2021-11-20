#!/usr/bin/env bash

function yell () { >&2 echo "$*";  }
function die () { yell "$*"; exit 1; }
function try () { "$@" || die "Command failed: $*"; }

dry_run=false

regex='PackageReference Include="([^"]*)" Version="([^"]*)"'
find . -name "*.*proj" | while read proj
do
    while read line
    do
        if [[ $line =~ $regex ]]; then
            name="${BASH_REMATCH[1]}"
            version="${BASH_REMATCH[2]}"
            if [[ $version != *-* ]]; then
                if $dry_run; then
                    echo "dotnet add $proj package $name"
                else
                    try dotnet add $proj package $name
                fi
            fi
        fi
    done < $proj
done
