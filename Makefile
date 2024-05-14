.PHONY: init reset swagger-gen protoc api grpc cdktf grpc-build api-build

init:
	@go work init
	@go work use api grpc rmq-consumer cdktf protos
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@export PATH="$PATH:$(go env GOPATH)/bin"


reset:
	@mv ./cdktf/generated ./generated
	@killall gopls
	@sleep 2
	@mv ./generated ./cdktf

swagger-gen:
	@cd ./api && swag init && cd ../

protoc:
	@protoc \
		--go_out=./protos --go_opt=module=github.com/ralvescosta/ec2-hellow-world/protos \
    --go-grpc_out=./protos --go-grpc_opt=module=github.com/ralvescosta/ec2-hellow-world/protos \
		--experimental_allow_proto3_optional \
    ./protos/def/service.proto

api:
	@cd api
	@GO_ENV=local go run ./api/main.go api

grpc:
	@cd grpc
	@GO_ENV=local go run ./grpc/main.go grpc

cdktf:
	@cd cdktf
	@GO_ENV=staging cdktf plan

grpc-build:
	@cd grpc
	@docker build . -t rafaelbodao/ecs-grpc:latest --network=host
	@docker push rafaelbodao/ecs-grpc:latest

api-build:
	@cd api
	@docker build . -t rafaelbodao/ecs-api:latest
	@docker push rafaelbodao/ecs-api:latest
