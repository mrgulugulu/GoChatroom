#!/bin/bash
#cd $WORKSPACE
export GOPROXY=https://goproxy.io
 
 #根据 go.mod 文件来处理依赖关系。
go mod tidy
 
# linux环境编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
 
# 构建docker镜像，项目中需要在当前目录下有dockerfile，否则构建失败

docker build -t chatserver .
docker tag  chatserver hermanyep/chatserver





