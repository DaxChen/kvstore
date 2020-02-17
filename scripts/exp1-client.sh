#!/bin/bash

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


echo ""
echo "=============================="
echo "=== LOAD DATA              ==="
echo "=============================="
echo ""
sleep 1
./kvclient load $num_keys $value_size
echo ""
echo "=== DONE LOAD DATA         ==="
echo ""


echo ""
echo "=============================="
echo "=== EXP1 READ TEST         ==="
echo "=============================="
echo ""
sleep 1
./kvclient exp1 read $num_keys
echo ""
echo "=== DONE EXP1 READ TEST    ==="
echo ""


echo ""
echo "=============================="
echo "=== EXP1 READ/WRITE TEST   ==="
echo "=============================="
echo ""
sleep 1
./kvclient exp1 readwrite $num_keys $value_size
echo ""
echo "=== DONE EXP1 READ/WRITE   ==="
echo ""
