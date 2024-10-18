REM
REM Licensed to the TerrAPI under one or more contributor
REM license agreements. See the NOTICE file
REM distributed with this work for additional information
REM regarding copyright ownership. The TerrAPI licenses this file
REM to you under the Apache License, Version 2.0 (the
REM "License"); you may not use this file except in compliance
REM with the License.  You may obtain a copy of the License at
REM
REM     http://www.apache.org/licenses/LICENSE-2.0
REM
REM Unless required by applicable law or agreed to in writing, software
REM distributed under the License is distributed on an "AS IS" BASIS,
REM WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
REM See the License for the specific language governing permissions and
REM limitations under the License.
REM

@echo off
setlocal

rem Define array elements
set "PROTO_NAMES=activity deployment health"

rem Loop through each element in the array
for %%i in (%PROTO_NAMES%) do (
    protoc %%i/%%i.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

    if ERRORLEVEL 1 (
        echo error processing %%i.proto
        exit /b %ERRORLEVEL%
    )
)

endlocal