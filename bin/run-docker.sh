#!/bin/sh

APP_NAME="${PWD##*/}"

docker run \
    -p 3000:3000 \
    "$APP_NAME":latest
