init:
	@go work init
	@go work use api grpc rmq-consumer cdktf protos
	@go install github.com/swaggo/swag/cmd/swag@latest

reset:
	@mv ./cdktf/generated ../../generated
	@killall gopls
	@slep 2
	@mv ../../generated ./cdktf/generated

swagger-gen:
	@cd ./api && swag init && cd ../

protoc:
	@echo "protoc"

run-api:
	@GO_ENV=local go run ./api/main.go api