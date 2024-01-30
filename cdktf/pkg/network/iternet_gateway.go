package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/internetgateway"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewInternetGateway(stack *stack.MyStack) {
	internetGatewayName := fmt.Sprintf("%v-igw", stack.Cfgs.AppName)
	stack.InternetGateway = internetgateway.NewInternetGateway(stack.TfStack, jsii.String(internetGatewayName), &internetgateway.InternetGatewayConfig{
		VpcId: stack.Vpc.Id(),
	})
}
