#!/usr/bin/env bash

AUTOMATE_DIRECTORY=/home/brad/automate/automate_core/
LOCK_FILE=$AUTOMATE_DIRECTORY/data/lock.json

echo {"isLocked":false} > $LOCK_FILE


