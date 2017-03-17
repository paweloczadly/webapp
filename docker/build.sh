#!/bin/bash
# cleanup old data
rm -rf public webapp

cd ..
cp -r public docker/public
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/webapp
cd docker
docker-compose build