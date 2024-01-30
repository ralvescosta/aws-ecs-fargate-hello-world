package containers

import "github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"

func NewEcsContainers(stack *stack.MyStack) {
	NewNginxContainer(stack)
}
