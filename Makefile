protos:
	protoc -I protos protos/example.proto --go_out=plugins=grpc:.

.PHONY: protos