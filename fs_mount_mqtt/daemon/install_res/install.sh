#!/usr/bin/env bash
sudo mkdir -p /opt/automate/
sudo cp -r automate_core /opt/automate
sudo cp automated /etc/init.d
sudo systemctl daemon-reload
