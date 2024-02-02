package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/eip"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/natgateway"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewNatGateway(stack *stack.MyStack) {
	eipName := fmt.Sprintf("%v-nat-g-eip", stack.Cfgs.AppName)
	eip := eip.NewEip(stack.TfStack, jsii.String(eipName), &eip.EipConfig{
		Domain: jsii.String("vpc"),
		// Instance: stack.NatGateway.Id(),
	})

	natGatewayName := fmt.Sprintf("%v-nat-g", stack.Cfgs.AppName)
	stack.NatGateway = natgateway.NewNatGateway(stack.TfStack, jsii.String(natGatewayName), &natgateway.NatGatewayConfig{
		SubnetId:         stack.PrivateSubnet.Id(),
		ConnectivityType: jsii.String("public"),
		AllocationId:     eip.Id(),
	})

	return
}
