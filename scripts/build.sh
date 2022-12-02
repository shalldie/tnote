#!/bin/bash

set -e

BASE_PATH=$(
    cd $(dirname $0)/..
    pwd
)

cd $BASE_PATH

export CGO_ENABLED=0

TARGET_OS_NAMES=(linux-amd64 linux-arm64 darwin-amd64 darwin-arm64)

for os_name in ${TARGET_OS_NAMES[*]}; do

    tupleName=(${os_name//-/ })

    echo build $os_name ...

    GOOS=${tupleName[0]} \
        GOARCH=${tupleName[1]} \
        go build -gcflags "all=-trimpath=$BASE_PATH" -o ttm.${os_name} main.go

    mkdir -p output
    mv ttm.$os_name output/ttm.$os_name

done
