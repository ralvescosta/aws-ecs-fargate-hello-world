package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewSubnets(stack *stack.MyStack) {
	privateSubNetName := fmt.Sprintf("%v-private-sbnt", stack.Cfgs.AppName)
	stack.PrivateSubnet = subnet.NewSubnet(stack.TfStack, jsii.Sprintf(privateSubNetName), &subnet.SubnetConfig{
		VpcId:            stack.Vpc.Id(),
		CidrBlock:        jsii.String(stack.Cfgs.PrivateSubnetCIDR),
		AvailabilityZone: jsii.String(stack.Cfgs.PrivateSubnetAZ),
	})

	publicSubNetName := fmt.Sprintf("%v-public-sbnt", stack.Cfgs.AppName)
	stack.PublicSubnet = subnet.NewSubnet(stack.TfStack, jsii.Sprintf(publicSubNetName), &subnet.SubnetConfig{
		VpcId:            stack.Vpc.Id(),
		CidrBlock:        jsii.Sprintf(stack.Cfgs.PublicSubnetCIDR),
		AvailabilityZone: jsii.String(stack.Cfgs.PublicSubnetAZ),
	})
}
