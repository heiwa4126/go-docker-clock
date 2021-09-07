#!/bin/sh
# see: <https://github.com/golang/go/wiki/WindowsCrossCompiling>
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-docker-clock -ldflags="-w -s" -trimpath main.go
upx --best go-docker-clock
