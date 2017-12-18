#!/usr/bin/env bash

mkdir $(pwd)/go
GOPATH="$(pwd)/go"
export GOPATH
go get github.com/influxdata/influxdb/client/v2
