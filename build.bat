@echo off
echo Building MonoMind for all platforms...

REM Create dist directory if it doesn't exist
if not exist "dist" mkdir "dist"

REM Build for Linux (AMD64)
echo Building for Linux (AMD64)...
set GOOS=linux
set GOARCH=amd64
go build -o dist\mono-linux-amd64 ./cmd/mono

REM Build for Windows (AMD64)
echo Building for Windows (AMD64)...
set GOOS=windows
set GOARCH=amd64
go build -o dist\mono-windows-amd64.exe ./cmd/mono

REM Build for macOS (AMD64)
echo Building for macOS (AMD64)...
set GOOS=darwin
set GOARCH=amd64
go build -o dist\mono-darwin-amd64 ./cmd/mono

REM Build for macOS (ARM64 - Apple Silicon)
echo Building for macOS (ARM64)...
set GOOS=darwin
set GOARCH=arm64
go build -o dist\mono-darwin-arm64 ./cmd/mono

echo Build complete! Binaries are located in the dist\ directory:
echo   - Linux (AMD64): dist\mono-linux-amd64
echo   - Windows (AMD64): dist\mono-windows-amd64.exe
echo   - macOS (AMD64): dist\mono-darwin-amd64
echo   - macOS (ARM64): dist\mono-darwin-arm64

echo All binaries have been built successfully!
pause