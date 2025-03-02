#!/bin/bash
set -e

echo "Building Rust backend..."
cargo build --release

echo "Rust backend built successfully!"
