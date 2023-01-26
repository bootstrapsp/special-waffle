protos:
	protoc -I protos protos/example.proto --go-grpc_out=. --go_out=.

.PHONY: protos