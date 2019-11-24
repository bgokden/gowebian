#!/bin/sh
path=$1
GOOS=js GOARCH=wasm go build -o $path/main.wasm $path/main.go
