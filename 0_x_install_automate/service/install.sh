#!/usr/bin/env bash
sudo mkdir -p /opt/automated/
sudo cp start.sh /opt/automated
sudo cp stop.sh /opt/automated
sudo cp status.sh /opt/automated
sudo cp use_image.sh /opt/automated

(
   cd /opt/automated/
   ./use_image.sh 'bradfordmedeiros/automate_arm_0.7.0'
)


sudo cp automated /etc/init.d
sudo systemctl daemon-reload
sudo service automated start
