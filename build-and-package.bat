@echo off
echo Building and packaging MonoMind for all platforms...

REM Clean previous builds
if exist "dist" rmdir /s /q "dist"
if exist "release" rmdir /s /q "release"

REM Create directories
mkdir "dist"
mkdir "release"

REM Build for all platforms
echo Building binaries...
call build.bat

REM Copy binaries to release directory
echo Copying binaries to release directory...
copy "dist\*" "release\" >nul

REM Rename binaries with version
echo Renaming binaries with version...
cd release
for /f %%i in (..\version.txt) do set VERSION=%%i
ren "mono-linux-amd64" "mono-%VERSION%-linux-amd64"
ren "mono-windows-amd64.exe" "mono-%VERSION%-windows-amd64.exe"
ren "mono-darwin-amd64" "mono-%VERSION%-darwin-amd64"
ren "mono-darwin-arm64" "mono-%VERSION%-darwin-arm64"

REM Create release index.html
echo Creating release index.html...
copy "..\release\index.html" "." >nul

REM Create README.md
echo Creating release README.md...
copy "..\release\README.md" "." >nul

echo.
echo Build and packaging complete!
echo Release files are available in the release/ directory:
dir
pause