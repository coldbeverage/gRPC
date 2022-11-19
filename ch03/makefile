.ONESHELL:

clean-server:
	rm -Rf ./ordermanagement/server/ecommerce

clean-client:
	rm -Rf ./ordermanagement/client/ecommerce

clean: clean-server clean-client
	rm -Rf bin
	mkdir bin

generate-server-messages:
	protoc --go_out=./ordermanagement/server ecommerce/order_management.proto

generate-server-services:
	protoc --go-grpc_out=./ordermanagement/server ecommerce/order_management.proto

generate-server: clean-server generate-server-messages generate-server-services

generate-client-messages:
	protoc --go_out=./ordermanagement/client ecommerce/order_management.proto

generate-client-services:
	protoc --go-grpc_out=./ordermanagement/client ecommerce/order_management.proto

generate-client: clean-client generate-client-messages generate-client-services

generate: generate-server generate-client

build-server:
	cd ordermanagement/server
	go build -v -o ../../bin/server
	cd ../..

build-client:
	cd ordermanagement/client
	go build -v -o ../../bin/client
	cd ../..