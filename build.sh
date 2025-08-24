#!/bin/bash

# Build script for MonoMind - Cross-platform compilation
# This script builds binaries for Windows, macOS, and Linux

echo "Building MonoMind for all platforms..."

# Create dist directory if it doesn't exist
mkdir -p dist

# Build for Linux (AMD64)
echo "Building for Linux (AMD64)..."
GOOS=linux GOARCH=amd64 go build -o dist/mono-linux-amd64 ./cmd/mono

# Build for Windows (AMD64)
echo "Building for Windows (AMD64)..."
GOOS=windows GOARCH=amd64 go build -o dist/mono-windows-amd64.exe ./cmd/mono

# Build for macOS (AMD64)
echo "Building for macOS (AMD64)..."
GOOS=darwin GOARCH=amd64 go build -o dist/mono-darwin-amd64 ./cmd/mono

# Build for macOS (ARM64 - Apple Silicon)
echo "Building for macOS (ARM64)..."
GOOS=darwin GOARCH=arm64 go build -o dist/mono-darwin-arm64 ./cmd/mono

echo "Build complete! Binaries are located in the dist/ directory:"
echo "  - Linux (AMD64): dist/mono-linux-amd64"
echo "  - Windows (AMD64): dist/mono-windows-amd64.exe"
echo "  - macOS (AMD64): dist/mono-darwin-amd64"
echo "  - macOS (ARM64): dist/mono-darwin-arm64"

# Make Linux and macOS binaries executable
chmod +x dist/mono-linux-amd64
chmod +x dist/mono-darwin-amd64
chmod +x dist/mono-darwin-arm64

echo "Linux and macOS binaries have been made executable."