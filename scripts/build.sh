#!/bin/bash

set -e

BASE_PATH=$(
    cd $(dirname $0)/..
    pwd
)

cd $BASE_PATH

export CGO_ENABLED=0

# darwin-arm64 不能使用 upx 压缩
TARGET_OS_NAMES=(linux-amd64 linux-arm64 darwin-amd64)

mkdir -p output

for os_name in ${TARGET_OS_NAMES[*]}; do

    tupleName=(${os_name//-/ })

    echo build $os_name ...

    GOOS=${tupleName[0]} \
        GOARCH=${tupleName[1]} \
        go build \
        -gcflags "all=-trimpath=$BASE_PATH" \
        -ldflags="-s -w" \
        -o output/tnote.${os_name} cmd/main.go

done

upx output/tnote.*
