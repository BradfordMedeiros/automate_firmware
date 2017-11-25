#!/usr/bin/env bash

mkdir $(pwd)/go
GOPATH="$(pwd)/go"
export GOPATH
go get github.com/eclipse/paho.mqtt.golang
go get golang.org/x/net/websocket
go get golang.org/x/net/proxy

