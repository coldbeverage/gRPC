
clean:
	rm -Rf ./productinfo/service/ecommerce

generate-server:
	protoc --go_out=./productinfo/service --go-grpc_out=./productinfo/service ecommerce/product_info.proto

generate: clean generate-server