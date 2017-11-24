#!/usr/bin/env bash

mkdir -p ./build
go build -o ./build/fs_mount_mqttd -ldflags "-s -w" ./src/*
cp ./install_res/install.sh ./build/install.sh
cp ./install_res/uninstall.sh ./build/uninstall.sh
cp ./install_res/fs_mount_mqtt ./build/fs_mount_mqtt