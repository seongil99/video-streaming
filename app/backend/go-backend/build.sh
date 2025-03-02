#!/bin/bash
set -e

# Create bin directory if it doesn't exist
mkdir -p bin

# Build the Go application
echo "Building Go backend..."
go build -o bin/go-backend main.go

echo "Go backend built successfully!"
