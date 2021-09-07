#!/bin/sh
# see: <https://github.com/golang/go/wiki/WindowsCrossCompiling>
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o clock -ldflags="-w -s" clock.go
