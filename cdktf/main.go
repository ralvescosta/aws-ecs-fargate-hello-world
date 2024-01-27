package main

import (
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func main() {
	cfgs := configs.NewConfigs()

	appScope, tfScope := pkg.NewAWSScopeProvider(cfgs)
	pkg.ApplyStack(cfgs, tfScope)

	appScope.Synth()
}
