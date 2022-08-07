#!/bin/bash
go get -u google.golang.org/protobuf/proto
brew install protoc-gen-go
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$GOPATH/bin
protoc --go_out=. *.proto