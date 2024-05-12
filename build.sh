#!/bin/bash

docker run --rm -it -v "$PWD":/_ -w /_ golang:1-alpine sh -c '
    go get -d ./...

    dirname="dist"
    filename="epgstation-slack-notification"

    (GOOS=darwin GOARCH=arm64 go build -o "${dirname}/darwin/${filename}") &
    (GOOS=linux GOARCH=arm GOARM=7 go build -o "${dirname}/linux-arm-7/${filename}") &
    wait
'
