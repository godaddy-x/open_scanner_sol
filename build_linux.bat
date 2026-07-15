@echo off
echo ==========================================
echo   [1/5] Setting environment to Linux ARM64...
echo ==========================================
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64

echo Cleaning old build artifacts...
if exist "output" rmdir /s /q output
mkdir output

echo ==========================================
echo   [2/5] Compiling SOL-Scanner Program...
echo ==========================================
go build -installsuffix cgo -ldflags="-s -w" -o output/scanner-sol main.go
if %errorlevel% neq 0 (
    echo [ERROR] SOL-Scanner compilation failed!
    goto :restore_env
)

echo ==========================================
echo   [3/5] Compressing SOL-Scanner with UPX...
echo ==========================================
upx --best --lzma output/scanner-sol

:restore_env
echo ==========================================
echo   [4/5] Restoring environment to Windows...
echo ==========================================
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64

echo ==========================================
echo   [5/5] Done! Binaries are in the 'output' folder.
echo ==========================================
pause
