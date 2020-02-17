#!/bin/bash

rm -f mem.log

echo ""
echo "=============================="
echo "=== STARTING SERVER        ==="
echo "=============================="
echo ""
sleep 1
./kvserver &
SERVER_PID=$!
sleep 3

server_addr="localhost:10000"
num_keys=10000
value_size=512
num_clients=$((10))

echo ""
echo "=============================="
echo "=== LOAD DATA              ==="
echo "=============================="
echo ""
sleep 1
./kvclient -server_addr=$server_addr -command=load -num_keys=$num_keys -value_size=$value_size
echo ""
echo "=== DONE LOAD DATA         ==="
echo ""


echo ""
echo "=============================="
echo "=== SPAWNING 10 CLIENTS    ==="
echo "=============================="
echo ""
last_client_pid=0
for i in $(seq 1 $num_clients); do
  echo "spawning client $i"
  sleep 1
  ./kvclient -server_addr=$server_addr -command=exp1 -mode=readwrite -num_keys=$num_keys -value_size=$value_size -duration=20 &
  last_client_pid=$!
done
wait $last_client_pid
echo ""
echo "=== DONE 10 CLIENTS        ==="
echo ""

echo ""
echo "=============================="
echo "=== GET STAT               ==="
echo "=============================="
echo ""
sleep 1
./kvclient -server_addr=$server_addr -command=stat
echo ""
echo "=== DONE GET STAT          ==="
echo ""


echo "killing server..."
kill $SERVER_PID
