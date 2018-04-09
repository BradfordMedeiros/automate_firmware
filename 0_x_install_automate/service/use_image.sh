#!/usr/bin/env bash

image=$1

docker pull $1
echo "$1" > /opt/automated/container.config
