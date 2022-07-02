#!/bin/bash
rm -f star_atlas_server
# go mod tidy
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
docker build --platform=linux/amd64 -t star_atlas_server .
