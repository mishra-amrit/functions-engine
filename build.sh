#!/usr/bin/env bash
rm -rf ./bin
mkdir -p bin/configuration bin/data/log bin/deployments
cp *config.json bin/configuration/

source ./env-setup.sh
go build -o bin/myappengine src/myappengine.io/appengine.go
chmod a+x bin/myappengine