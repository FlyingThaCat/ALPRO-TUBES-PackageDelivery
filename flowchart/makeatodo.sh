#!/bin/bash
set -euo pipefail

# Navigate to project root
cd "$(dirname "$0")/.." || exit 1

# Prepare output paths
output="flowchart/go_functions.md"
mkdir -p flowchart
> "$output"

# Counter for indexing
index=1

find . -type f -name "*.go" ! -path "*/.git/*" | while read -r file; do
  lineno=0
  filename=$(basename "$file")
  while IFS= read -r line; do
    lineno=$((lineno + 1))
    if echo "$line" | grep -qE '^[[:space:]]*func[[:space:]]+(\([^)]+\)[[:space:]]*)?[A-Za-z_][A-Za-z0-9_]*[[:space:]]*\('; then
      # Extract function name
      name=$(echo "$line" | sed -E 's/^[[:space:]]*func[[:space:]]+(\([^)]+\)[[:space:]]*)?([A-Za-z_][A-Za-z0-9_]*).*/\2/')
      
      # Write to MD with index and checkbox
      echo "$index # $file:$lineno  # $name # [ ]" >> "$output"
      
      # Create corresponding empty D2 file
      touch "flowchart/${index}.${name}.d2"

      index=$((index + 1))
    fi
  done < "$file"
done

echo "Done. Markdown written to $output and D2 files created in flowchart/"
