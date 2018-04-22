#!/usr/bin/env bash
sudo service automated stop
sudo rm -rf /opt/automated
sudo rm -rf /etc/init.d/automated
sudo systemctl disable automated-controller
sudo docker kill automate-controller
sudo docker rm automate-controller
sudo rm /lib/systemd/system/automated-controller.service
sudo systemctl daemon-reload
