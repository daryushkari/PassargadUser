GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto: proto/user.proto
	rm -rf api
	mkdir -p api/pb
	protoc --go_out=api/pb --go_opt=paths=source_relative \
        --go-grpc_out=api/pb --go-grpc_opt=paths=source_relative \
        proto/user.proto

.PHONY: config
config:
	cp ./secret-example.json ./secret-config.json
