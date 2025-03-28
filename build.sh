#!/bin/bash

# 1. 获取程序名。
DIR_NAME=$(basename $(pwd))
VER=$(TZ="Europe/London" date -d "$(git log -1 --format=%cd --date=iso)"  +"%y.%-m.%-d-%H%M%S")
BIN_NAME=${DIR_NAME}"-"v$VER$(go env GOEXE)
echo "程序名为："${BIN_NAME}

# 2. 如果执行的是清理命令，清理完就退出。
go clean -cache
if [ "$1" = "clean" ]; then
    rm -rf ${DIR_NAME}*
    echo "已清理"
    exit 0
fi

export CGO_ENABLED=0
LDFLAGS="-s -w -extldflags -static"
go build -o ${BIN_NAME} -trimpath -v -ldflags "$LDFLAGS" ./main

echo "编译打包结束"
