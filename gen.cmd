@echo off
setlocal

rem Define array elements
set "PROTO_NAMES=activity deployment"

rem Loop through each element in the array
for %%i in (%PROTO_NAMES%) do (
    protoc %%i/%%i.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

    if ERRORLEVEL 1 (
        echo error processing %%i.proto
        exit /b %ERRORLEVEL%
    )
)

endlocal