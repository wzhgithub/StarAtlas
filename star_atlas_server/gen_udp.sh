#!/bin/bash
# 第一个0是vmc的id 第二个是连的交换机id
./utest /dev/stdout 0 0 | xargs --null echo > /dev/udp/127.0.0.1/9191