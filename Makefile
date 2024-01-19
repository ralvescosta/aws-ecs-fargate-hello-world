init:
	go work init
	go work use api grpc rmq-consumer cdktf protos

reset:
	mv ./cdktf/generated ../../generated
	killall gopls
	slep 2
	mv ../../generated ./cdktf/generated

protoc: