#!/bin/bash

# Build untuk Windows
GOOS=windows GOARCH=amd64 go build -o bin/windows/loopshpere.exe main.go

# Build untuk macOS
GOOS=darwin GOARCH=amd64 go build -o bin/macos/loopshpere main.go

# Build untuk Linux
GOOS=linux GOARCH=amd64 go build -o bin/linux/loopshpere main.go

echo "Build selesai! Executable tersedia di folder bin."