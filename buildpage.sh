#!/bin/sh
path=$1
# project=$2
mkdir -p $path/public
go generate $path/main.go
GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o $path/public/main.wasm $path/main.go
gzip --best -k -f $path/public/main.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js $path/public/

# tinygo build -o $path/main.wasm -target wasm $path/main.go

# docker run -v $GOPATH:/go -v $(pwd):/go/src/${project} -e "GOPATH=/go" tinygo/tinygo:0.9.0 tinygo build -o /go/src/${project}/${path}/main.wasm -target wasm --no-debug /go/src/${project}/${path}/main.go
