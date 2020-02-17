#!/bin/bash

num_keys=$((1 * 1024 * 1024))
value_size=$((4 * 1024))

load_data() {
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
}

# ask user if want to load data
echo "Do you want to load 4GB data?"
select yn in "Yes" "No"; do
    case $yn in
        Yes ) load_data; break;;
        No ) break;;
    esac
done


# get number of clients from user
read -e -p "Number of clients: " num_clients
if [ $num_clients -lt 1 ]; then
  echo "ERROR input! number of clients should be positive"
  exit 1
fi

echo ""
echo "=============================="
echo "=== EXP3 READ TEST         ==="
echo "=============================="
echo ""
sleep 1
for i in {1..$num_clients}
do
  ./kvclient exp1 read $num_keys
done
echo ""
echo "=== DONE EXP3 READ TEST    ==="
echo ""


echo ""
echo "=============================="
echo "=== EXP3 READ/WRITE TEST   ==="
echo "=============================="
echo ""
sleep 1
./kvclient exp1 readwrite $num_keys $value_size
echo ""
echo "=== DONE EXP3 READ/WRITE   ==="
echo ""
