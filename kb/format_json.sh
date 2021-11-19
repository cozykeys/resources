#!/usr/bin/env bash

function format_json() {
    local json_file="$1"

    local tmp_file="$(mktemp)"
    if [ ! "$?" = "0" ]; then
        >&2 echo "mktemp failed"
        return 1
    fi

    jq "$json_file" > "$tmp_file"
    if [ ! "$?" = "0" ]; then
        >&2 echo "jq failed"
        return 1
    fi

    cat "$tmp_file"

    rm "$tmp_file"
}
