@echo off

:: Build untuk Windows
set GOOS=windows
set GOARCH=amd64
go build -o bin\windows\apirostream.exe main.go

:: Build untuk macOS
set GOOS=darwin
set GOARCH=amd64
go build -o bin\macos\apirostream main.go

:: Build untuk Linux
set GOOS=linux
set GOARCH=amd64
go build -o bin\linux\apirostream main.go

echo Build selesai! Executable tersedia di folder bin.