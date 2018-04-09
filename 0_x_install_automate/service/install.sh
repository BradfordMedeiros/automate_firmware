#!/usr/bin/env bash
sudo mkdir -p /opt/automated/
sudo cp start.sh /opt/automated
sudo cp stop.sh /opt/automated
sudo cp status.sh /opt/automated

sudo cp automated /etc/init.d
sudo systemctl daemon-reload
sudo service automated start
