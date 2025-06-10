#!/bin/bash
set -euo pipefail

FONT="Noto Sans Mono, Noto Color Emoji 14"
FG_COLOR="white"
BG_COLOR="black"

# Find all .dirty files recursively
find . -type f -name "*.dirty" | while IFS= read -r dirtyfile; do
    echo "Processing $dirtyfile"

    # Define output paths
    cleanfile="${dirtyfile%.dirty}.log"
    imagefile="${dirtyfile%.dirty}.png"

    # Clean ANSI escape codes
    perl -pe 's/\e\[[0-9;]*[A-Za-z]//g' "$dirtyfile" > "$cleanfile"

    # Render to image using pango-view
    pango-view \
        --no-display \
        --font="$FONT" \
        --foreground="$FG_COLOR" \
        --background="$BG_COLOR" \
        --output="$imagefile" \
        "$cleanfile"

    # Remove the original dirty file
    rm "$dirtyfile"
    rm "$cleanfile"
    echo "Created $imagefile and cleaned up $dirtyfile"
done
