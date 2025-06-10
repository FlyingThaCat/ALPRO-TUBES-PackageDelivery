#!/bin/bash
set -euo pipefail

output_dir="output"

mkdir -p "$output_dir"

for file in *.d2; do
  filename=$(basename "$file" .d2)
  d2 "$file" "$output_dir/${filename}.png"
done

echo "✅ All diagrams compiled to $output_dir/"
