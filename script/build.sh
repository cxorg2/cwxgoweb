#!/bin/bash

work_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $work_dir/../


rm bin/*

export CGO_ENABLED=0
go build -o bin/cwxgoweb -ldflags '-s -w' src/*.go
