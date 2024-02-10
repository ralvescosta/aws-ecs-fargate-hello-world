package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/eip"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/natgateway"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewNatGateway(stack *stack.MyStack) {
	eipAName := fmt.Sprintf("%v-nat-g-eip-a", stack.Cfgs.AppName)
	eipA := eip.NewEip(stack.TfStack, jsii.String(eipAName), &eip.EipConfig{
		Domain: jsii.String("vpc"),
		// Instance: stack.NatGateway.Id(),
	})

	natGatewayAName := fmt.Sprintf("%v-nat-g-a", stack.Cfgs.AppName)
	stack.NatGateways.PrivateA = natgateway.NewNatGateway(stack.TfStack, jsii.String(natGatewayAName), &natgateway.NatGatewayConfig{
		SubnetId:         stack.Subnets.PublicA.Id(),
		ConnectivityType: jsii.String("public"),
		AllocationId:     eipA.Id(),
	})

	eipBName := fmt.Sprintf("%v-nat-g-eip-b", stack.Cfgs.AppName)
	eipB := eip.NewEip(stack.TfStack, jsii.String(eipBName), &eip.EipConfig{
		Domain: jsii.String("vpc"),
		// Instance: stack.NatGateway.Id(),
	})

	natGatewayBName := fmt.Sprintf("%v-nat-g-b", stack.Cfgs.AppName)
	stack.NatGateways.PrivateB = natgateway.NewNatGateway(stack.TfStack, jsii.String(natGatewayBName), &natgateway.NatGatewayConfig{
		SubnetId:         stack.Subnets.PublicB.Id(),
		ConnectivityType: jsii.String("public"),
		AllocationId:     eipB.Id(),
	})

	return
}
