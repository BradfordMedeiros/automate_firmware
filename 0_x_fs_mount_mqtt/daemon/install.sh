#!/usr/bin/env bash

mkdir $HOME/.go
GOPATH="$HOME/.go"
export GOPATH
go get github.com/eclipse/paho.mqtt.golang
go get golang.org/x/net/websocket
go get golang.org/x/net/proxy
go get github.com/nu7hatch/gouuid

