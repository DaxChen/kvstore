#!/bin/bash
cd server
go build -o main.exe
cd ..
cd client
go build -o main.exe
cd ..

############### experiment 1 ###############
#echo "experiment 1"
#echo "clean mem.log and build server"
#rm ./mem.log
#$"./server/main.exe" &
#
#sleep 10
#
#echo "512B write work load"
#$"./client/main.exe" "exp1" "512B.txt"
#echo "512B write read load"
#$"./client/main.exe" "exp1" "keys.txt"
#echo "512B write r/w load"
#$"./client/main.exe" "exp1" "mix512B.txt"
#
#echo "4KB write work load"
#$"./client/main.exe" "exp1" "4KB.txt"
#echo "4KB write read load"
#$"./client/main.exe" "exp1" "keys.txt"
#echo "4KB write r/w load"
#$"./client/main.exe" "exp1" "mix4KB.txt"
#
#echo "512KB write work load"
#$"./client/main.exe" "exp1" "512KB.txt"
#echo "512KB write read load"
#$"./client/main.exe" "exp1" "keys.txt"
#echo "512KB write r/w load"
#$"./client/main.exe" "exp1" "mix512KB.txt"
#
#echo "1MB write work load"
#$"./client/main.exe" "exp1" "1MB.txt"
#echo "1MB write read load"
#$"./client/main.exe" "exp1" "keys.txt"
#echo "1MB write r/w load"
#$"./client/main.exe" "exp1" "mix1MB.txt"
#
#echo "4MB write work load"
#$"./client/main.exe" "exp1" "4MB.txt"
#echo "4MB write read load"
#$"./client/main.exe" "exp1" "keys.txt"
#echo "4MB write r/w load"
#$"./client/main.exe" "exp1" "mix4MB.txt"
#
#kill -SIGTERM $(ps aux | pgrep -f './server/main.exe')





############### experiment 2 ###############
#echo "experiment 2"
#echo "clean mem.log and build server"
#rm ./mem.log
#$"./server/main.exe" &
#
#sleep 10
#
#echo "build 512B data"
#$"./client/main.exe" "exp1" "512B.txt"
#
## crash server
#$"./crashrecoverserver.sh" &
#
#echo "client spin"
#$"./client/main.exe" "exp2" "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
#
#kill -SIGTERM $(ps aux | pgrep -f './server/main.exe')





############### experiment 3 ###############
#echo "experiment 3"
#echo "clean mem.log and build server"
#rm ./mem.log
#$"./server/main.exe" &
#
#echo "2 clients write load"
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt"
#echo "2 clients write read load"
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt"
#echo "2 clients write r/w load"
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt"
#
#echo "4 clients write load"
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt"
#echo "4 clients write read load"
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt"
#echo "4 clients write r/w load"
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt"
#
#echo "8 clients write load"
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt"
#echo "8 clients write read load"
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt"
#echo "8 clients write r/w load"
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt"
#
#echo "16 clients write load"
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt" &
#$"./client/main.exe" "exp1" "512KB.txt"
#echo "16 clients write read load"
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt" &
#$"./client/main.exe" "exp1" "keys.txt"
#echo "16 clients write r/w load"
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt" &
#$"./client/main.exe" "exp1" "mix512KB.txt"
#
#kill -SIGTERM $(ps aux | pgrep -f './server/main.exe')

############### prefix & validate correctness ###############
echo "prefix"
echo "clean mem.log and build server"
rm ./mem.log
$"./server/main.exe" &

echo "set prefix data"
$"./client/main.exe" "exp1" "prefix.txt"
echo "check prefix"
$"./client/main.exe" "prefix" "A"

kill -SIGTERM $(ps aux | pgrep -f './server/main.exe')

echo "recover server"
$"./server/main.exe" &
echo "check prefix again"
$"./client/main.exe" "prefix" "A"

kill -SIGTERM $(ps aux | pgrep -f './server/main.exe')
