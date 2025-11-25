#!/bin/bash

set -e

SCRIPT_DIR="$(dirname "$0")"
CSV_FILE="$SCRIPT_DIR/quotes.csv"
DRY_RUN=false

# Parse arguments
if [ "$1" = "--dry-run" ]; then
    DRY_RUN=true
    echo "=== DRY RUN MODE - No database operations will be performed ==="
    echo ""
fi

if [ ! -f "$CSV_FILE" ]; then
    echo "Error: quotes.csv not found at $CSV_FILE"
    exit 1
fi

# Use Python to properly parse CSV and generate SQL
read -r -d '' PYTHON_SCRIPT << 'EOF' || true
import csv
import sys

csv_file = sys.argv[1]
dry_run = sys.argv[2] == "true"

author_values = []
quote_values = []

with open(csv_file, 'r', encoding='utf-8') as f:
    reader = csv.reader(f)
    for row in reader:
        if len(row) != 3:
            continue

        quote_id = row[0].strip()
        message = row[1].strip()
        author_name = row[2].strip()

        if not quote_id:
            continue

        if dry_run:
            print(f"ID: {quote_id}")
            print(f"Quote: {message}")
            print(f"Author: {author_name}")
            print("-" * 80)
        else:
            # Escape single quotes for SQL
            message_escaped = message.replace("'", "''")
            author_escaped = author_name.replace("'", "''")

            author_values.append(f"({quote_id}, '{author_escaped}')")
            quote_values.append(f"({quote_id}, '{message_escaped}', {quote_id})")

if not dry_run:
    print("AUTHORS|||" + ",".join(author_values))
    print("QUOTES|||" + ",".join(quote_values))
EOF

# Run Python script
output=$(python3 -c "$PYTHON_SCRIPT" "$CSV_FILE" "$DRY_RUN")

if [ "$DRY_RUN" = true ]; then
    echo "$output"
    echo ""
    echo "=== DRY RUN COMPLETE - Review the parsed data above ==="
    exit 0
fi

# Extract SQL values from Python output
author_values=$(echo "$output" | grep "^AUTHORS|||" | cut -d'|' -f4)
quote_values=$(echo "$output" | grep "^QUOTES|||" | cut -d'|' -f4)

if [ -z "$author_values" ] || [ -z "$quote_values" ]; then
    echo "Error: Failed to parse CSV file"
    exit 1
fi

# Insert authors
echo "Inserting authors into author-service database..."
docker exec author-service-mysql mysql -uroot -proot author_service -e \
    "INSERT INTO authors (id, name) VALUES $author_values ON DUPLICATE KEY UPDATE name=VALUES(name);"

# Insert quotes
echo "Inserting quotes into quote-service database..."
docker exec quote-service-mysql mysql -uroot -proot quote_service -e \
    "INSERT INTO quotes (id, message, author_id) VALUES $quote_values ON DUPLICATE KEY UPDATE message=VALUES(message), author_id=VALUES(author_id);"

echo "Done! Loaded $(echo "$author_values" | grep -o '(' | wc -l) quotes and authors successfully."
