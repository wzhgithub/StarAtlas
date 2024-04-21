#!/bin/bash
docker rm -f $(docker ps -a -q)
sh build.sh 
docker volume rm star_atlas_server_data
docker compose up -d
