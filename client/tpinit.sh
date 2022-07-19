#!/bin/bash

for ((i=0;i<3;i++)); do 
  ./utest /dev/null ${i} 0
  sleep 1 
done;


