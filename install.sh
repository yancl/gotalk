#!/usr/bin/env bash

if [ ! -f install.sh ]; then
echo 'install must be run within its container folder' 1>&2
exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

#clean libs
rm -rf pkg/linux_amd64/*

gofmt -w src

go install searcher

#mv bin/main bin/apns_pusher

export GOPATH="$OLDGOPATH"

echo 'finished'
