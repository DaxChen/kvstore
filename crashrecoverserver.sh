#!/bin/bash
t="5"
echo "sleep $t seconds before crash server"
sleep $t
echo "crash"
kill -SIGTERM $(ps aux | pgrep -f './server/main.exe') # ps aux | grep -f './server/main.exe' | awk '{print $2}'
echo "restart"
$"./server/main.exe"
