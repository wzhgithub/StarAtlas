#!/bin/bash
docker rm -f $(docker ps -a -q)
sh build.sh 
docker compose up -d
