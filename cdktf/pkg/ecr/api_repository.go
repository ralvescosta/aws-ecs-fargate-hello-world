package ecr

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecrrepository"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewAPIECRRepository(stack *stack.MyStack) {
	repoName := fmt.Sprintf("%v-api", stack.Cfgs.AppName)
	stack.EcrAPIRepository = ecrrepository.NewEcrRepository(stack.TfStack, jsii.String(repoName), &ecrrepository.EcrRepositoryConfig{
		Name: jsii.String(repoName),
		ImageScanningConfiguration: &ecrrepository.EcrRepositoryImageScanningConfiguration{
			ScanOnPush: jsii.Bool(true),
		},
	})
}
