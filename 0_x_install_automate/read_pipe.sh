#!/bin/bash

touch /somefile
pipe=/opt/automated/pipe

trap "rm -f $pipe" EXIT

if [[ ! -p $pipe ]]; then
    mkfifo $pipe
fi

while true
do
    if read line <$pipe; then
        if [[ "$line" == 'quit' ]]; then
            break
        fi
        echo "line is: $line"
	eval $line
    fi
done

echo "Reader exiting"
