#!/usr/bin/env bash
set -e
echo "building linux executable"
GOOS=linux go build
docker build -t nhvtran/zipsvr .
docker push nhvtran/zipsvr
go clean