@ECHO OFF
setlocal
set "SCRIPTS_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPTS_DIR%\.."
set "BUILD_DIR=%PROJECT_DIR%\build"
cd "%PROJECT_DIR%\cmd\app"
go build -o "%BUILD_DIR%" && "%BUILD_DIR%\app.exe"
cd "%PROJECT_DIR%"
endlocal
