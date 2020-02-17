#!/bin/bash

# get server address from user
read -e -p "Server Address host:port (eg. localhost:10000)" server_addr

num_keys=$((1 * 1024 * 1024))
value_size=$((4 * 1024))
echo "num_keys: " $num_keys ", value_size: " $value_size

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
echo "=== EXP2 COLD TEST         ==="
echo "=============================="
echo ""
sleep 1
./kvclient -server_addr=$server_addr -command=exp2
echo ""
echo "=== DONE EXP2 COLD TEST    ==="
echo ""
