#! /bin/sh

docker run -d -p 9000:9000 -p 1883:1883 bradfordmedeiros/automate_arm_0.7.0

echo "starting" >> /opt/automated/log.txt
