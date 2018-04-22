#!/usr/bin/env bash

image=$1

docker pull $1
/opt/automated/stop-controller.sh
echo "$1" > /opt/automated/controller-container.config
/opt/automated/start-controller.sh
