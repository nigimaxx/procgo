#!/bin/sh

cd client
go build -o procgo

cd ../daemon
go build -o procgo-daemon

cd ..

cwd=$(pwd)

ln -s "$cwd/client/procgo" /usr/local/bin
ln -s "$cwd/daemon/procgo-daemon" /usr/local/bin

procgo completion
ln -s "$cwd/_procgo" /usr/local/share/zsh/site-functions

sudo mkdir /var/log/procgo
sudo touch /var/log/procgo/daemon.log
sudo chmod 666 /var/log/procgo/daemon.log
