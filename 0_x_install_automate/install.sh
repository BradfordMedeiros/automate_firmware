#!/usr/bin/env bash

echo "installing core docker image"

echo "copying files and scripts"
sudo mkdir -p /opt/automated/
sudo cp start.sh /opt/automated
sudo cp start-controller.sh /opt/automated
sudo cp stop.sh /opt/automated
sudo cp status.sh /opt/automated
sudo cp use_image.sh /opt/automated

sudo cp read_pipe.sh /opt/automated
sudo cp write_pipe.sh /opt/automated

echo "creating named pipe in /opt/automated/pipe"
sudo mkfifo /opt/automated/pipe  # main docker controller writes to this pipe, script started in service container reads constantly from pipe and executes it

(
   cd /opt/automated/
   echo "bradfordmedeiros/core_arm_1.0.0" > /opt/automated/container.config
)

echo "installing automated_controller service"
#sudo cp automated-controller /etc/init.d
sudo cp automated-controller /lib/systemd/system/automated-controller.service
sudo systemctl daemon-reload
sudo systemctl enable automated-controller
sudo service automated-controller start
echo "finished installing automated_controller service"

echo "installing automated service"
#sudo cp automated /etc/init.d
#sudo systemctl daemon-reload 
#sudo service automated start
echo "finished installing automated service"

echo "finished installing automate core docker image"
