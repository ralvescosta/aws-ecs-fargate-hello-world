package pkg

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/ecs"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/ecs/containers"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/network"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
	"go.uber.org/zap"
)

func ApplyStack(logger *zap.SugaredLogger, cfgs *configs.Configs, tfStack cdktf.TerraformStack) {
	myStack := stack.MyStack{
		Cfgs:    cfgs,
		Logger:  logger,
		TfStack: tfStack,
	}

	network.NewVpc(&myStack)
	network.NewSubnets(&myStack)
	network.NewInternetGateway(&myStack)
	network.NewNatGateway(&myStack)
	network.NewRouteTables(&myStack)
	network.NewApplicationLoadBalancer(&myStack)
	ecs.NewECSFargateCluster(&myStack)
	containers.NewEcsContainers(&myStack)
}
