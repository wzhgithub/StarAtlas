#!/bin/bash

if [[ ! -d "bin" ]]; then
  mkdir -p bin
fi

docker run -e USER="$(id -u)" -u="$(id -u)"  -v `pwd`:/app -w /app --platform linux/amd64 buildessential/debian:latest bash build.sh
