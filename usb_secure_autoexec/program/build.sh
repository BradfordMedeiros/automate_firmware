
#!/bin/bash

AUTOMATE_KEY=$1
USB_FILE_DIRECTORY=$2

if [ $# != 2 ]
  then
    echo "Invalid number of arguments to build usb" >&2
    echo "Usage: build.sh <AUTOMATE_KEY> <USB_FILE_DIRECTORY>" >&2
    exit 1
  else 
    echo "Using: AUTOMATE_KEY: $AUTOMATE_KEY; USB_FILE_DIRECTORY: $USB_FILE_DIRECTORY"
fi

mkdir ./build

cargo build --release

cp target/release/usb_secure_autoexec ./build
cp ./run_scripts.sh ./build

sed -i "s/<SCRIPT_PATH>/$USB_FILE_DIRECTORY/g" ./build/run_scripts.sh
sed -i "s/<SYS_PATH>/$AUTOMATE_KEY/g" ./build/run_scripts.sh

