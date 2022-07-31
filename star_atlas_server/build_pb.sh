#!/bin/bash
# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$GOPATH/bin
protoc --go_out=. *.proto