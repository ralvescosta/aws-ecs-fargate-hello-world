init:
	@go work init
	@go work use api grpc rmq-consumer cdktf protos
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@export PATH="$PATH:$(go env GOPATH)/bin"


reset:
	@mv ./cdktf/generated ../../generated
	@killall gopls
	@slep 2
	@mv ../../generated ./cdktf/generated

swagger-gen:
	@cd ./api && swag init && cd ../

protoc:
	@protoc \
		--go_out=./protos --go_opt=module=github.com/ralvescosta/ec2-hellow-world/protos \
    --go-grpc_out=./protos --go-grpc_opt=module=github.com/ralvescosta/ec2-hellow-world/protos \
		--experimental_allow_proto3_optional \
    ./protos/def/service.proto

run-api:
	@GO_ENV=local go run ./api/main.go api

run-grpc:
	@GO_ENV=local go run ./grpc/main.go grpc