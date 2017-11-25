#!/usr/bin/env bash
sudo service fs_mount_mqtt stop
sudo rm -rf /opt/fs_mount_mqtt
sudo rm -rf /etc/init.d/fs_mount_mqtt
sudo systemctl daemon-reload