#!/bin/sh
base_path=$1

mkdir -p $base_path/public
go generate $base_path/main.go
GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o $base_path/public/main.wasm $base_path/main.go
gzip --best -k -f $base_path/public/main.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js $base_path/public/