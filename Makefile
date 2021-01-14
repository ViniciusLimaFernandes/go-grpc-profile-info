protosrc = $(wildcard proto/*.proto)

proto: $(protosrc)
	protoc --go_out=plugins=grpc:. proto/user.proto

build:
	@echo "---- Building Application ----"
	@go build -o server-bin server/*.go
	@go build -o client-bin client/*.go

run:
	@echo "---- Running Server ----"
	@go run server/*.go

run_client:
	@echo "---- Running Client ----"
	@go run client/*.go