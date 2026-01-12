@echo off
echo Building Financial Toolbox Desktop Application...

echo.
echo Step 1: Cleaning build directory...
if exist build rmdir /s /q build
if exist dist rmdir /s /q dist
mkdir build

echo.
echo Step 2: Building frontend...
cd frontend
call npm install
call npm run build
cd ..

echo.
echo Step 3: Building Wails application...
wails build -clean -platform windows/amd64

echo.
echo Build completed!
echo Output: build\bin\financial-toolbox.exe
pause
