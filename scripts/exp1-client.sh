#!/bin/bash

# get server address from user
read -e -p "Server Address host:port (eg. localhost:10000)" server_addr

# get number of keys from user
read -e -p "Number of keys: " num_keys
if [ $num_keys -lt 1 ]; then
  echo "ERROR input! number of keys should be positive"
  exit 1
fi

# get value size from user
read -e -p "value size in bytes: " value_size
if [ $value_size -lt 1 ]; then
  echo "ERROR input! value size should be positive"
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
echo "=== EXP1 READ TEST         ==="
echo "=============================="
echo ""
sleep 1
./kvclient -server_addr=$server_addr -command=exp1 -mode=read -num_keys=$num_keys
echo ""
echo "=== DONE EXP1 READ TEST    ==="
echo ""


echo ""
echo "=============================="
echo "=== EXP1 READ/WRITE TEST   ==="
echo "=============================="
echo ""
sleep 1
./kvclient -server_addr=$server_addr -command=exp1 -mode=readwrite -num_keys=$num_keys -value_size=$value_size
echo ""
echo "=== DONE EXP1 READ/WRITE   ==="
echo ""
