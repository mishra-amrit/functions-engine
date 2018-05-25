#!/usr/bin/env bash
rm -rf ./bin
mkdir -p bin/configuration bin/data/log bin/deployments
cp *config.yaml bin/configuration/

source ./env-setup.sh
go build -o bin/myappengine src/myappengine.io/main.go
chmod a+x bin/myappengine