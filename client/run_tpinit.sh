#!/bin/bash

if [[ ! -f "bin/utest" ]]; then
  bash run.sh
fi

docker run -e USER="$(id -u)" -u="$(id -u)"  -v `pwd`:/app -w /app --platform linux/amd64 --network host buildessential/debian:latest bash tpinit.sh
