#!/bin/bash

echo "Please answer if you want to remove mem.log:"
rm -i mem.log

echo "Starting server. Will Monitor and restart the server if crashes..."
until ./kvserver; do
    echo "Server './kvserver' crashed with exit code $?.  Respawning.." >&2
    # echo "DEBUG mode: respawning in one second to avoid fast crash..."
    # sleep 1
done
