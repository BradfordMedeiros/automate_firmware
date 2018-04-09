#! /bin/sh

echo "about to stop" >> /opt/automated/log.txt
docker stop automate
docker rm automate


echo "finished stopping" >> /opt/automated/log.txt
