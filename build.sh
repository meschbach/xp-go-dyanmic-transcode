#!/bin/zsh

./gen-grpc.sh
go build -o client ./client
go build -o service
