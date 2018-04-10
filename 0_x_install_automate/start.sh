#! /bin/bash

container=$(</opt/automated/container.config)
echo "container is $container"

docker run -d -p 9000:9000 -p 1883:1883 --name=automate "$container"

echo "starting" >> /opt/automated/log.txt
