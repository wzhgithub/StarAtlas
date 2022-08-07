#!/bin/bash

#for ((i=0;i<3;i++)); do 
#  ./utest /dev/null ${i} 0
#  sleep 1 
#done;


./utest /dev/null 1 3
sleep 1
./utest /dev/null 2 3
sleep 1

./utest /dev/null 3 1
sleep 1
./utest /dev/null 4 1
sleep 1
./utest /dev/null 5 1
sleep 1

./utest /dev/null 6 6
