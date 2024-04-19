#!/bin/bash
set -x
docker rm -f $(docker ps -a -q)
docker volume rm star_atlas_server_data
bash build.sh
docker compose -f docker-compose-no-fe.yaml up -d
