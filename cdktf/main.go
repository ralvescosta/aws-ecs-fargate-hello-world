package main

import (
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
	"go.uber.org/zap"
)

func main() {
	logger := zap.S().Named("cdktf")

	cfgs := configs.NewConfigs(logger)

	appScope, tfScope := pkg.NewAWSScopeProvider(cfgs)
	pkg.ApplyStack(logger, cfgs, tfScope)

	appScope.Synth()
}
