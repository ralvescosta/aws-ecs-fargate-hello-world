package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/eip"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/natgateway"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// This function will create two nat gateways
//
// The NAT Gateway will be responsible to configure the internet Egress for the private subnets
// it mens that, the private subnet will be able to have outbound traffic to internet
func NewNatGateway(stack *stack.MyStack) {
	eipAName := fmt.Sprintf("%v-nat-g-eip-a", stack.Cfgs.AppName)
	stack.NatGateways.EIpA = eip.NewEip(stack.TfStack, jsii.String(eipAName), &eip.EipConfig{
		Domain: jsii.String("vpc"),
		// Instance: stack.NatGateway.Id(),
	})

	natGatewayAName := fmt.Sprintf("%v-nat-g-a", stack.Cfgs.AppName)
	stack.NatGateways.PrivateA = natgateway.NewNatGateway(stack.TfStack, jsii.String(natGatewayAName), &natgateway.NatGatewayConfig{
		SubnetId:         stack.Subnets.PublicA.Id(),
		ConnectivityType: jsii.String("public"),
		AllocationId:     stack.NatGateways.EIpA.Id(),
	})

	eipBName := fmt.Sprintf("%v-nat-g-eip-b", stack.Cfgs.AppName)
	stack.NatGateways.EIpB = eip.NewEip(stack.TfStack, jsii.String(eipBName), &eip.EipConfig{
		Domain: jsii.String("vpc"),
		// Instance: stack.NatGateway.Id(),
	})

	natGatewayBName := fmt.Sprintf("%v-nat-g-b", stack.Cfgs.AppName)
	stack.NatGateways.PrivateB = natgateway.NewNatGateway(stack.TfStack, jsii.String(natGatewayBName), &natgateway.NatGatewayConfig{
		SubnetId:         stack.Subnets.PublicB.Id(),
		ConnectivityType: jsii.String("public"),
		AllocationId:     stack.NatGateways.EIpB.Id(),
	})

	return
}
