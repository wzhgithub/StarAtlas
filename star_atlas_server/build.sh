#!/bin/bash
go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
docker build -t star_atlas_server .
