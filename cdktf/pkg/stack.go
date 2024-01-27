package pkg

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/ecs"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/network"
)

func ApplyStack(cfgs *configs.Configs, tfStack cdktf.TerraformStack) {
	fnaVpc := network.NewVpc(cfgs, tfStack)
	privateA, privateB, publicA, publicB := network.NewSubnets(cfgs, tfStack, fnaVpc)
	igw := network.NewInternetGateway(cfgs, tfStack, fnaVpc)
	natGtwA, natGtwB := network.NewNatGateway(cfgs, tfStack, publicA, publicB)
	_, _, _ = network.NewRouteTables(cfgs, tfStack, fnaVpc, igw, privateA, privateB, publicA, publicB, natGtwA, natGtwB)
	albSecGroup, albTargetGroup, fnaAlb := network.NewApplicationLoadBalancer(cfgs, tfStack, fnaVpc, publicA, publicB)
	ecs.NewECSFargate(cfgs, tfStack, fnaVpc, privateA, privateB, albTargetGroup, albSecGroup, fnaAlb)
}
