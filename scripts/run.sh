#!/bin/bash

# to start from root directory
cd "$(dirname "$0")" || exit

# check if go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

echo "Running the application..."
