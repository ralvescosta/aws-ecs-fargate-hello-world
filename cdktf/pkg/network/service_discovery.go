package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/servicediscoveryprivatednsnamespace"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// This method will create the Discovery Private Namespace
//
// Private Namespace allow ECS Services to communicate between than
func NewServiceDiscoveryPrivateNamespace(stack *stack.MyStack) {
	privateDNSNamespaceName := fmt.Sprintf("%v-private-dns-namespace", stack.Cfgs.AppName)
	stack.ServiceDiscoveryPrivateNamespace = servicediscoveryprivatednsnamespace.NewServiceDiscoveryPrivateDnsNamespace(stack.TfStack, jsii.String(privateDNSNamespaceName), &servicediscoveryprivatednsnamespace.ServiceDiscoveryPrivateDnsNamespaceConfig{
		Name:        jsii.String(stack.Cfgs.AppName),
		Vpc:         stack.Vpc.Id(),
		Description: jsii.String("Namespace for internal ECS services communication"),
	})
}
