package ecs

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscluster"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// The method will create ECS cluster
func NewECSFargateCluster(stack *stack.MyStack) {
	ecsClusterName := fmt.Sprintf("%v-ecs-cluster", stack.Cfgs.AppName)

	stack.EcsCluster = ecscluster.NewEcsCluster(stack.TfStack, jsii.String(ecsClusterName), &ecscluster.EcsClusterConfig{
		Name: jsii.String(ecsClusterName),
		Setting: []*ecscluster.EcsClusterSetting{
			{
				Name:  jsii.String("containerInsights"),
				Value: jsii.String("enabled"),
			},
		},
	})
}
