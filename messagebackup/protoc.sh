#!/bin/bash

# https://github.com/aurora-is-near/devops-stuff/tree/main/docker/protoc
docker run --rm -v "$(pwd)"/:/proto protoc --proto_path=/proto/ --go_out=/proto/. --go_opt=paths=source_relative *.proto 
