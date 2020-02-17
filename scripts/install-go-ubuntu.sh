#!/bin/bash

# install go
echo "Downloading and installing Go from official website..."
curl -O https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz
tar -xvf go1.13.8.linux-amd64.tar.gz
sudo chown -R root:root ./go
sudo mv go /usr/local

# setup path
echo 'export GOPATH=$HOME/go' >> ~/.profile
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.profile
source ~/.profile

# verify
go version
