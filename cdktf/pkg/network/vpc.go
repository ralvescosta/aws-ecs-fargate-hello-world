package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// This method will create the VPC
//
// VPC is the internal network that all services will be configured in it
func NewVpc(stack *stack.MyStack) {
	vpcName := fmt.Sprintf("%v-vpc", stack.Cfgs.AppName)

	stack.Vpc = vpc.NewVpc(stack.TfStack, jsii.String(vpcName), &vpc.VpcConfig{
		CidrBlock: jsii.Sprintf(stack.Cfgs.VpcCIDR),
	})
}
