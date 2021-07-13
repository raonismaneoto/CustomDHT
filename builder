#!/bin/bash

readonly grpc_dir=./commons/grpc_api
readonly grpc_api_file=$grpc_dir/api.pb.go

[ ! -f $grpc_api_file ] && cd $grpc_dir && protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./api.proto && cd -
