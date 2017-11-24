#!/usr/bin/env bash

(cd ./daemon/ && ./clean.sh)
(cd ./client/ && ./clean.sh)

rm -r ./build/