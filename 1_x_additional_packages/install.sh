#!/usr/bin/env bash

echo "installing additional firmware"

# json command line parsing
sudo apt-get install jq

# hobbiest community loves node (maybe not all... lol dur dur javscript haters)
wget -O - https://raw.githubusercontent.com/sdesalas/node-pi-zero/master/install-node-v9.8.0.sh | bash

# nice package to be able to publish/subscribe to topics from command line
npm install -g mqtt

ln -s /opt/nodejs/bin/mqtt /usr/local/bin/mqtt

echo "finsihed installing additional firmware"
