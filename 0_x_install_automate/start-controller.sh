#! /bin/bash



cd /opt/automated/
./read_pipe.sh &


container=$(</opt/automated/controller-container.config)

docker rm automate-controller
docker run -d -v /opt/automated/:/opt/automated/ -p 9999:9999 --name=automate-controller "$container"

echo "starting" >> /opt/automated/log.txt
