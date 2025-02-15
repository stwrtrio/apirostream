@echo off

:: Build untuk Windows
set GOOS=windows
set GOARCH=amd64
go build -o bin\windows\loopshpere.exe main.go

:: Build untuk macOS
set GOOS=darwin
set GOARCH=amd64
go build -o bin\macos\loopshpere main.go

:: Build untuk Linux
set GOOS=linux
set GOARCH=amd64
go build -o bin\linux\loopshpere main.go

echo Build selesai! Executable tersedia di folder bin.