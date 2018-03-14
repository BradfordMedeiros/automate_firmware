
#!/bin/bash

function sedeasy {
  sed -i "s/$(echo $1 | sed -e 's/\([[\/.*]\|\]\)/\\&/g')/$(echo $2 | sed -e 's/[\/&]/\\&/g')/g" $3
}

AUTOMATE_KEY=$1             # path where they automate hash be
USB_FILE_DIRECTORY=$2       # usb file directory to execute scripts from, we assume the key is automate.key on the root (maybe way to use udev to do this instead?)
EXECUTABLE_PATH=$3          # path where the usb_secure_autoexec executable will be put


if [ $# != 3 ]
  then
    echo "Invalid number of arguments to build usb" >&2
    echo "Usage: build.sh <AUTOMATE_KEY> <USB_FILE_DIRECTORY> <EXECUTABLE_PATH>" >&2
    exit 1
  else 
    echo "Using: AUTOMATE_KEY: $AUTOMATE_KEY; USB_FILE_DIRECTORY: $USB_FILE_DIRECTORY"
fi

mkdir ./build

cargo build --release

cp target/release/usb_secure_autoexec ./build
cp ./run_scripts.sh ./build

sedeasy "<SCRIPT_PATH>" "$USB_FILE_DIRECTORY" ./build/run_scripts.sh
sedeasy "<SYS_PATH>" "$AUTOMATE_KEY" ./build/run_scripts.sh

cp ./udev/10-autorun.rules ./build
sedeasy "<EXECUTABLE>" "$EXECUTABLE_PATH" ./build/10-autorun.rules
