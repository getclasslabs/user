#!/bin/bash

echo "Compiling the API"
docker run -it --rm -v "$(pwd)":/go -e GOPATH= golang:1.14 sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a --installsuffix cgo --ldflags='-s'"

rm ./docker/user
mv ./user ./docker/
cp ./config.yaml ./docker/

docker build -t getclass/user:$1 docker/

docker push getclass/user:$1

if [[ ! $(docker service ls | grep gcl_user) = "" ]]; then
  docker service update gcl_user --image getclass/user:$1
else
  docker stack deploy -c docker-compose.yaml gcl
fi