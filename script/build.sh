#!/bin/bash

work_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $work_dir/../

pwd

rm -f bin/*

export CGO_ENABLED=0
go build -o bin/cwxgoweb -ldflags '-s -w' cmd/*.go

# go build -o bin/cwxgoweb -ldflags '-linkmode "external" -extldflags "-static" -s -w' cmd/*.go
