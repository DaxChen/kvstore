#!/bin/bash

# get server address from user
read -e -p "Server Address host:port (eg. localhost:10000)" server_addr

num_keys=$((1 * 1024 * 1024))
value_size=$((4 * 1024))

# get number of clients from user
read -e -p "Number of clients: " num_clients
if [ $num_clients -lt 1 ]; then
  echo "ERROR input! number of clients should be positive"
  exit 1
fi

load_data() {
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
}

# ask user if want to load data
echo "Do you want to load data? num_keys=$num_keys, value_size=$value_size"
select yn in "Yes" "No"; do
    case $yn in
        Yes ) load_data; break;;
        No ) break;;
    esac
done


echo ""
echo "=============================="
echo "=== EXP3 READ TEST         ==="
echo "=============================="
echo ""
sleep 1
for i in $(seq 1 $num_clients); do
  echo "spawning client $i"
  ./kvclient -server_addr=$server_addr -command=exp1 -mode=read -num_keys=$num_keys &
done
wait
echo ""
echo "=== DONE EXP3 READ TEST    ==="
echo ""


echo ""
echo "=============================="
echo "=== EXP3 READ/WRITE TEST   ==="
echo "=============================="
echo ""
sleep 1
for i in $(seq 1 $num_clients); do
  echo "spawning client $i"
  ./kvclient -server_addr=$server_addr -command=exp1 -mode=readwrite -num_keys=$num_keys -value_size=$value_size &
done
wait
echo ""
echo "=== DONE EXP3 READ/WRITE   ==="
echo ""
