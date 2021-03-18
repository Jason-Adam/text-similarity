#!/bin/sh

APP_NAME="${PWD##*/}"
IMG_TAG="$(git rev-parse --short HEAD)"
IMG=docker.io/jasonadam/"$APP_NAME"

docker build \
    -t "$IMG":"$IMG_TAG" \
    -t "$IMG":latest .

docker push "$IMG":latest
docker push "$IMG":"$IMG_TAG"
