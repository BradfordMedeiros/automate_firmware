#!/usr/bin/env bash

mkdir ./build

(cd ./daemon/ && source ./install.sh && ./build.sh)
cp ./daemon/build/* ./build/

(cd ./client/ && source ./install.sh && ./build.sh)
cp ./client/build/* ./build