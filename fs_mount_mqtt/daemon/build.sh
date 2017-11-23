#!/usr/bin/env bash

mkdir -p ./build
go build -o ./build/fs_mount_mqttd -ldflags "-s -w" ./src/*
