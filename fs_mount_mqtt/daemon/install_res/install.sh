#!/usr/bin/env bash
sudo mkdir -p /opt/fs_mount_mqtt/
sudo cp -r fs_mount_mqttd /opt/fs_mount_mqtt
sudo cp fs_mount_mqtt /etc/init.d
sudo systemctl daemon-reload
