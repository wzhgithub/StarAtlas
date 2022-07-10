#!/bin/bash

for ((i=0;i<3;i++)); do 
  bin/utest /dev/stdout ${i} 0 | xargs --null echo > /dev/udp/127.0.0.1/9191 
  sleep 1 
done;


