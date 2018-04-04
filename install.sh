#!/usr/bin/env bash

# This file is responsible for invoking build all the necessary projects and then copying them into the correct places
# This may eventually become a proper make file at some point

# Basically automate is a raspberry pi, put inside a docker container, with an SSH service exposed outside of the docker container so we
# can upgrade the pi.  Volumes mapped to store persistant data + allow gpio access.

# Install docker onto this machine
curl -sSL https://get.docker.com | sh
