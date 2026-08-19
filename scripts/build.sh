#!/bin/bash

set -e

BASE_PATH=$(
    cd $(dirname $0)/..
    pwd
)

cd $BASE_PATH

export CGO_ENABLED=0

# 注意：darwin-arm64 不能使用 upx 压缩（Apple Silicon 的 Mach-O 压缩后会损坏），
# 因此它只构建、不压缩，见下方 upx 步骤中的跳过逻辑。
TARGET_OS_NAMES=(linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64.exe)

# 清空旧产物，避免 upx 重复压缩已压过的文件（AlreadyPackedException）
rm -rf output
mkdir -p output

for os_name in ${TARGET_OS_NAMES[*]}; do

    # 先剥离可能存在的后缀（如 .exe），再按 - 拆分出 GOOS/GOARCH
    tupleName=(${os_name%%.*})
    tupleName=(${tupleName//-/ })

    echo build $os_name ...

    GOOS=${tupleName[0]} \
        GOARCH=${tupleName[1]} \
        go build \
        -gcflags "all=-trimpath=$BASE_PATH" \
        -ldflags="-s -w" \
        -o output/tnote.${os_name} main.go

done

# 压缩除 darwin-arm64 之外的所有产物（upx 无法正确压缩 Apple Silicon 二进制）
# !(...) 为 bash 扩展通配，需先开启 extglob
shopt -s extglob
upx --force-macos output/tnote.!(*darwin-arm64)
