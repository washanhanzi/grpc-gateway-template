.PHONY: genproto dev-gateway dev

genproto:
	@protoc --proto_path proto --go_out=plugins=grpc:. ./proto/example.proto
	@protoc --proto_path proto --grpc-gateway_out=logtostderr=true:. ./proto/example.proto 

dev-gateway: export GATEWAY_BINDING_ADDR=localhost:9192
dev-gateway:
	@go run gateway/main.go

dev:
	@go run main.go