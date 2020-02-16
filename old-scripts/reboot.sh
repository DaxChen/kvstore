#!/bin/bash
t="10"
echo "sleep $t seconds before reboot server"
sleep $t

echo "reboot"
$"./server/main.exe"
