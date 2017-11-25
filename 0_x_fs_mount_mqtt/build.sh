#!/usr/bin/env bash

mkdir ./build
mkdir -p ./build/client
mkdir -p ./build/daemon

(cd ./daemon/ && source ./install.sh && ./build.sh)
cp ./daemon/build/* ./build/daemon

(cd ./client/ && source ./install.sh && ./build.sh)
cp ./client/build/* ./build/client