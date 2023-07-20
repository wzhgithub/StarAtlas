#!/bin/bash
docker rm -f fe
docker run -d -p 80:80 --name fe fe_nginx
