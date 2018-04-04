#!/usr/bin/env bash

# This file is responsible for invoking build all the necessary projects and then copying them into the correct places
# This may eventually become a proper make file at some point

# Basically automate is a raspberry pi, put inside a docker container, with an SSH service exposed outside of the docker container so we
# can upgrade the pi.  Volumes mapped to store persistant data + allow gpio access.


### Install dependencies section ###

# Install docker onto this machine.  Is there a better way of doing this?
if [ -z "$(which docker)" ]; then
    echo "docker is not yet installed"
    #curl -sSL https://get.docker.com | sh
fi 

### Configuring software into right places section ###

sudo service ssh start
echo "pi:password" | sudo chpasswd  # hardcoding password to password (as opposed to pi's default password of raspberry).  This probably needs to be done manually per device before shipping and saved. 

sudo useradd automate
echo "automate:password" | sudo chpasswd
mkdir /home/automate

