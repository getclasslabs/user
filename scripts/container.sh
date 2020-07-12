#!/bin/bash

echo "Compiling the API"
docker run -it --rm -v "$(pwd)":/go -e GOPATH= golang:1.14 sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a --installsuffix cgo --ldflags='-s'"

mv ./user ./docker/
cp ./config.yaml ./docker/
docker build -t getclass/user:latest docker/

docker push getclass/user:latest

kubectl create -f deployment.yaml