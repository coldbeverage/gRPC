.ONESHELL:

clean-server:
	rm -Rf ./productinfo/service/ecommerce

clean-client:
	rm -Rf ./productinfo/client/ecommerce

clean: clean-server clean-client
	rm -Rf bin
	mkdir bin

generate-server-messages:
	protoc --go_out=./productinfo/server ecommerce/product_info.proto

generate-server-services:
	protoc --go-grpc_out=./productinfo/server ecommerce/product_info.proto

generate-server: clean-server generate-server-messages generate-server-services

generate-client-messages:
	protoc --go_out=./productinfo/client ecommerce/product_info.proto

generate-client-services:
	protoc --go-grpc_out=./productinfo/client ecommerce/product_info.proto

generate-client: clean-client generate-client-messages generate-client-services

generate: generate-server generate-client

build-server:
	cd productinfo/server
	go build -v -o ../../bin/server
	cd ../..

build-client:
	cd productinfo/client
	go build -v -o ../../bin/client
	cd ../..