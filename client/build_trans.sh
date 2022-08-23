#!/bin/bash

gcc -Isrc -Isrc/third_party/curl/include src/trans.cc src/third_party/curl/libs/libcurl.a -std=c++11 -lstdc++ -lpthread -o bin/trans
