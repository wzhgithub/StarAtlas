#!/bin/bash
set -ex
rm -fr ./log/*
cp go.mod.bk go.mod
go mod tidy
go build
# docker compose rm -f
# docker compose -f debug/docker-compose-local-mongo.yaml up -d
export REPORT_PORT=12322
export REPORT_HOST="0.0.0.0"
./star_atlas_server -log_dir=./log -alsologtostderr -path ./debug/config.yaml
