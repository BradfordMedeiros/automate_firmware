#!/usr/bin/env bash
sudo service automated stop
sudo rm -rf /opt/automated
sudo rm -rf /etc/init.d/automated
sudo systemctl daemon-reload
