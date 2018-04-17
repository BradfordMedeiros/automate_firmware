#!/usr/bin/env bash

echo "installing core docker image"

sudo mkdir -p /opt/automated/
sudo cp start.sh /opt/automated
sudo cp stop.sh /opt/automated
sudo cp status.sh /opt/automated
sudo cp use_image.sh /opt/automated

sudo cp read_pipe.sh /opt/automated
sudo cp write_pipe.sh /opt/automated

(
   cd /opt/automated/
   echo "bradfordmedeiros/core_arm_1.0.0" > /opt/automated/container.config
)

sudo cp automated /etc/init.d
sudo systemctl daemon-reload
sudo service automated start

echo "finished install core docker image"
