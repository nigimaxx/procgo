#!/bin/sh

cd client
go build -o procgo

cd ../daemon
go build -o procgo-daemon
