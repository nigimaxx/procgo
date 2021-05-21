#!/bin/sh

cd client
go build -o procgo

cd ../daemon
go build -o procgo-daemon

cd ..

cwd=$(pwd)

ln -s "$cwd/client/procgo" /usr/local/bin
ln -s "$cwd/daemon/procgo-daemon" /usr/local/bin
