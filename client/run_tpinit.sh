#!/bin/bash

docker run -e USER="$(id -u)" -u="$(id -u)"  -v `pwd`:/app -w /app --platform linux/amd64 buildessential/debian:latest bash tpinit.sh
