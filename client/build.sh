#!/bin/bash
gcc -Isrc -Isrc/third_party/rapidjson/include  -g src/client_new.cc src/proto/*.cc src/common/*.cc  -std=c++11 -lstdc++ -o bin/utest
