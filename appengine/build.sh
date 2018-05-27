#!/usr/bin/env bash

echo "Initializing build.."
rm -rf ./bin
echo "Creating directory structure.."
mkdir -p bin/configuration bin/data/log bin/deployments bin/runtime
echo "Copying configurations.."
cp *config.yaml bin/configuration/
echo "Copying runtimes.."
cp -rv src/myappengine.io/runtime/* bin/runtime

echo "Initializing environment variables"
source ./env-setup.sh
echo "Building.."
go build -o bin/myappengine src/myappengine.io/main.go
echo "Setting up execute permission to build output"
chmod a+x bin/myappengine
echo "Build complete !"
