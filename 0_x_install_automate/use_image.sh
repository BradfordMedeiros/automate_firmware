#!/usr/bin/env bash

image=$1

docker pull $1
/opt/automated/stop.sh
echo "$1" > /opt/automated/container.config
/opt/automated/start.sh
