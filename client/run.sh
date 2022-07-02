#!/bin/bash

docker run -e USER="$(id -u)" -u="$(id -u)"  -v `pwd`:/app -w /app buildessential/debian:latest bash build.sh
