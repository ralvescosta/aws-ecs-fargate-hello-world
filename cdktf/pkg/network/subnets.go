package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// This function will create 4 subnets under our VPC, two private and two publics
//
// The private subnets will be routed through the NAT Gateway and the public will be routed through the internet gateway
func NewSubnets(stack *stack.MyStack) {
	privateSubNetAName := fmt.Sprintf("%v-private-sbnt-a", stack.Cfgs.AppName)
	stack.Subnets.PrivateA = subnet.NewSubnet(stack.TfStack, jsii.Sprintf(privateSubNetAName), &subnet.SubnetConfig{
		VpcId:            stack.Vpc.Id(),
		CidrBlock:        jsii.String(stack.Cfgs.PrivateSubnetA_CIDR),
		AvailabilityZone: jsii.String(stack.Cfgs.PrivateSubnetA_AZ),
	})

	privateSubNetBName := fmt.Sprintf("%v-private-sbnt-b", stack.Cfgs.AppName)
	stack.Subnets.PrivateB = subnet.NewSubnet(stack.TfStack, jsii.Sprintf(privateSubNetBName), &subnet.SubnetConfig{
		VpcId:            stack.Vpc.Id(),
		CidrBlock:        jsii.String(stack.Cfgs.PrivateSubnetB_CIDR),
		AvailabilityZone: jsii.String(stack.Cfgs.PrivateSubnetB_AZ),
	})

	publicSubNetAName := fmt.Sprintf("%v-public-sbnt-a", stack.Cfgs.AppName)
	stack.Subnets.PublicA = subnet.NewSubnet(stack.TfStack, jsii.Sprintf(publicSubNetAName), &subnet.SubnetConfig{
		VpcId:            stack.Vpc.Id(),
		CidrBlock:        jsii.Sprintf(stack.Cfgs.PublicSubnetA_CIDR),
		AvailabilityZone: jsii.String(stack.Cfgs.PublicSubnetA_AZ),
	})

	publicSubNetBName := fmt.Sprintf("%v-public-sbnt-b", stack.Cfgs.AppName)
	stack.Subnets.PublicB = subnet.NewSubnet(stack.TfStack, jsii.Sprintf(publicSubNetBName), &subnet.SubnetConfig{
		VpcId:            stack.Vpc.Id(),
		CidrBlock:        jsii.Sprintf(stack.Cfgs.PublicSubnetB_CIDR),
		AvailabilityZone: jsii.String(stack.Cfgs.PublicSubnetB_AZ),
	})
}
