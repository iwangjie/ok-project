#!/usr/bin/env bash
version="2021.1.5"
name=${PWD##*/}
go test .
GOOS=linux go build -ldflags="-s -w" -o linux/"$name"

upx linux/"$name"

docker rmi -f --no-prune iwangjie/"$name":latest
docker build -t iwangjie/"$name":latest .
docker image tag iwangjie/"$name":latest 172.19.121.47:5000/iwangjie/"$name":latest
docker image push 172.19.121.47:5000/iwangjie/"$name":latest


docker rmi --no-prune iwangjie/"$name":"$version"
docker build -t iwangjie/"$name":"$version" .
docker image tag iwangjie/"$name":"$version" 172.19.121.47:5000/iwangjie/"$name":"$version"
docker image push 172.19.121.47:5000/iwangjie/"$name":"$version"

echo "推送镜像完成"