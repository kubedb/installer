#!/bin/bash

# File in which the replacements will be made
FILE="cassandra-pod.json"  # Replace with your file name

# Use sed to search and replace the format
sed -i 's/"legendFormat": *"\([^"]*\)"/"legendFormat": {{ `\1` }}/g' "$FILE"

echo "All legendFormat fields have been updated!"
