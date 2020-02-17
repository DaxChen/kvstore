#!/bin/bash

num_keys=$((1 * 1024 * 1024))
value_size=$((4 * 1024))
echo "num_keys: " $num_keys ", value_size: " $value_size

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
echo "=== EXP2 COLD TEST         ==="
echo "=============================="
echo ""
sleep 1
./kvclient exp2
echo ""
echo "=== DONE EXP2 COLD TEST    ==="
echo ""
