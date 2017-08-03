#!/bin/bash
source ./.env

docker build -t $appName .

docker stop "$appName" && docker rm -f "$appName"

docker run -d -t \
	--restart=always \
	-p "$appPort":80 \
    --net=crs_backend \
    --name=$appName \
    -v "$PWD"/site:/go/src/"$appPath" \
    -v "$PWD"/resources:"$resourcesPath" \
    -v "$PWD"/src:/go/src \
    -v "$PWD"/logs:/var/log/http \
    -w /go/src/"$appPath" \
    $appName
