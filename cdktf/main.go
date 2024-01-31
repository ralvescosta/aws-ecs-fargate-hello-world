package main

import (
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func main() {
	cfgs, logger := configs.NewConfigs()

	appScope, tfScope := pkg.NewAWSScopeProvider(logger, cfgs)
	pkg.ApplyStack(logger, cfgs, tfScope)

	appScope.Synth()
}
