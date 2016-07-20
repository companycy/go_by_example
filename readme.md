cmd to setup
	- gb vendor fetch golang.org/x/net/context
	- gb vendor fetch google.golang.org/grpc
	- wget https://github.com/google/protobuf/releases/download/v3.0.0-beta-4/protoc-3.0.0-beta-4-linux-x86_64.zip
	- protoc --go_out=plugins=grpc:pb ./requester.proto

cmd to build
	- gb build

cmd to run
	- bin/server
	- bin/client
