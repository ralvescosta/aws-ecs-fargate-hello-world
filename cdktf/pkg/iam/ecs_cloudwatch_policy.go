package iam

import (
	"fmt"
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/iampolicy"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/iamrole"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/iamrolepolicyattachment"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewIAMCloudWatch(stack *stack.MyStack) {
	ecsTaskExecutionRoleName := fmt.Sprintf("%v-ecsTaskExecutionRole", stack.Cfgs.AppName)
	stack.IAMCloudWatch.Role = iamrole.NewIamRole(stack.TfStack, jsii.String(ecsTaskExecutionRoleName), &iamrole.IamRoleConfig{
		AssumeRolePolicy: jsii.String(`{
					"Version": "2012-10-17",
					"Statement": [
							{
									"Effect": "Allow",
									"Principal": {"Service": "ecs-tasks.amazonaws.com"},
									"Action": "sts:AssumeRole"
							}
					]
			}`),
	})

	policy := `
	{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": [
					"logs:CreateLogStream",
					"logs:PutLogEvents"
				],
				"Resource": [
					"arn:aws:logs:<<AWS_REGION>>:<<ACCOUNT_ID>>:log-group:*"
				]
			}
		]
	}`

	policy = strings.Replace(policy, "<<AWS_REGION>>", stack.Cfgs.Region, -1)
	policy = strings.Replace(policy, "<<ACCOUNT_ID>>", stack.Cfgs.AccountID, -1)

	ecsLoggingPolicyName := fmt.Sprintf("%v-ecsLoggingPolicy", stack.Cfgs.AppName)
	stack.IAMCloudWatch.Policy = iampolicy.NewIamPolicy(stack.TfStack, jsii.String(ecsLoggingPolicyName), &iampolicy.IamPolicyConfig{
		Policy: jsii.String(policy),
	})

	ecsLoggingPolicyAttachmentName := fmt.Sprintf("%v-ecsLoggingPolicyAttachment", stack.Cfgs.AppName)
	stack.IAMCloudWatch.RolePolicyAtt = iamrolepolicyattachment.NewIamRolePolicyAttachment(stack.TfStack, jsii.String(ecsLoggingPolicyAttachmentName), &iamrolepolicyattachment.IamRolePolicyAttachmentConfig{
		Role:      stack.IAMCloudWatch.Role.Name(),
		PolicyArn: stack.IAMCloudWatch.Policy.Arn(),
	})
}
