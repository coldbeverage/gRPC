# gRPC
Repo For working through gRPC Up and Running


## Steps to configure to build/run with gRPC
The books steps for getting my environment up and running I believe is a bit outdated.  The recommended steps and compiler command I had to update to reflect what I think is more normal for this moment.

Running Ubuntu 22.04 on WSL, using Homebrew used the followed the steps provided by the gRPC [quickstart][grpc-go-quickstart] and that was much more effective.

```bash
# install protoc
> brew install protobuf
> protoc --version  #version 3+ is recommended
libprotoc 3.21.9

# install golang compiler plugins 
> go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Generator Error
When I started trying to generate the code based on the chapter 2 productinfo.proto file i ran into the following:

```bash
> protoc -I ecommerce ecommerce/product_info.proto --go_out=plugins=grpc:github.com/coldbeverage/productinfo/service/ecommerce
--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC

See https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code for more information.
```

via stack overflow the issue is due to changes release of a new set of tools to generate the code. This is the so called new way...  More details [here](https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable) 
```bash
  # generate the messages
 protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"

 # generate the services
 protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
 ```


## Additional Resources 
- [gRPC](https://grpc.io/)
- [formatting discussion for option go_package in .proto][stackoverflow-link]

[stackoverflow-link]: https://stackoverflow.com/questions/61666805/correct-format-of-protoc-go-package
[grpc-go-quickstart]: https://grpc.io/docs/languages/go/quickstart/