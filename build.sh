#!/bin/bash

# Build untuk Windows
GOOS=windows GOARCH=amd64 go build -o bin/windows/apirostream.exe main.go

# Build untuk macOS
GOOS=darwin GOARCH=amd64 go build -o bin/macos/apirostream main.go

# Build untuk Linux
GOOS=linux GOARCH=amd64 go build -o bin/linux/apirostream main.go

echo "Build selesai! Executable tersedia di folder bin."