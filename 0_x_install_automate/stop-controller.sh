#! /bin/sh

echo "about to stop" >> /opt/automated/log.txt
docker stop automate-controller
docker rm automate-controller


echo "finished stopping" >> /opt/automated/log.txt
