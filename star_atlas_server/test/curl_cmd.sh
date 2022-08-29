#!/bin/bash

#app 
curl -i -X POST -H "Content-Type: application/json;charset=utf-8" -d@failure_app.json \
"http://127.0.0.1:8088/vmc/do_failure_over"

curl -i -X POST -H "Content-Type: application/json;charset=utf-8" -d@failure_app.json \
"http://127.0.0.1:8088/vmc/do_failure_over"

#vmc
curl -i -X POST -H "Content-Type: application/json; charset=utf-8" -d@failure_vmc.json \
"http://127.0.0.1:8088/vmc/do_failure_over"

curl -i -X POST -H "Content-Type: application/json; charset=utf-8" -d@failure_vmc.json \
"http://127.0.0.1:8088/vmc/failure_over_result"