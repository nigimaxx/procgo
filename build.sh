#!/bin/sh

cd client
go build -o procgo

cd ../daemon
go build -o procgo-daemon

cd ..
procgo completion > /usr/local/share/zsh/site-functions/_procgo
