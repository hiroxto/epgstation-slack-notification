#!/bin/bash

docker run --rm -it -v "$PWD":/_ -w /_ golang:1-alpine sh -c '
    go get -d ./...

    dirname="dist"
    filename="epgstation-slack-notification"

    (GOOS=darwin GOARCH=arm64 go build -o "${dirname}/darwin-arm64/${filename}") &
    (GOOS=linux GOARCH=amd64 go build -o "${dirname}/linux-amd64/${filename}") &
    wait
'
