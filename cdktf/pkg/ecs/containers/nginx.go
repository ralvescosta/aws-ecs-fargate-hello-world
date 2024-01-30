package containers

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecsservice"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecstaskdefinition"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewNginxContainer(stack *stack.MyStack) {
	ecsNginxTaskDefinitionName := fmt.Sprintf("%v-ecs-nginx-td", stack.Cfgs.AppName)
	td := ecstaskdefinition.NewEcsTaskDefinition(stack.TfStack, jsii.String(ecsNginxTaskDefinitionName), &ecstaskdefinition.EcsTaskDefinitionConfig{
		Family:                  jsii.String("service"),
		Cpu:                     jsii.String("0.5"),
		Memory:                  jsii.String("128M"),
		NetworkMode:             jsii.String("awsvpc"),
		RequiresCompatibilities: jsii.Strings("FARGATE"),
		ContainerDefinitions: jsii.String(`
		[
			{
				"image": "nginx",
				"name": "fna-nginx",
				"portMappings": [{ "containerPort": 80 }]
			}
		]
		`),
	})

	ecsTaskDefinitionSecGroupName := fmt.Sprintf("%v-ecs-nginx-sec-group", stack.Cfgs.AppName)
	secGroup := securitygroup.NewSecurityGroup(stack.TfStack, jsii.String(ecsTaskDefinitionSecGroupName), &securitygroup.SecurityGroupConfig{
		Ingress: &[]*securitygroup.SecurityGroupIngress{
			{
				Protocol:       jsii.String("tcp"),
				FromPort:       jsii.Number(0),
				ToPort:         jsii.Number(6553),
				SecurityGroups: &[]*string{stack.ElasticLoadBalancerSecGroup.Id()},
			},
		},
		Egress: &[]*securitygroup.SecurityGroupEgress{
			{
				Protocol:   jsii.String("tcp"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				ToPort:     jsii.Number(0),
				FromPort:   jsii.Number(6553),
			},
		},
	})

	ecsServiceName := fmt.Sprintf("%v-ecs-nginx-svc", stack.Cfgs.AppName)
	ecsservice.NewEcsService(stack.TfStack, jsii.String(ecsServiceName), &ecsservice.EcsServiceConfig{
		Name:           jsii.String(ecsServiceName),
		Cluster:        stack.EcsCluster.Id(),
		TaskDefinition: td.Arn(),
		LaunchType:     jsii.String("FARGATE"),
		DesiredCount:   jsii.Number(2),
		DeploymentController: &ecsservice.EcsServiceDeploymentController{
			Type: jsii.String("ECS"),
		},
		NetworkConfiguration: &ecsservice.EcsServiceNetworkConfiguration{
			Subnets:        &[]*string{stack.PrivateSubnet.Id()},
			SecurityGroups: &[]*string{secGroup.Id()},
		},
		LoadBalancer: &[]*ecsservice.EcsServiceLoadBalancer{
			{
				ElbName:        stack.ElasticLoadBalancer.Name(),
				TargetGroupArn: stack.ElasticLoadBalancerSecGroup.Arn(),
				ContainerName:  jsii.String("fna-nginx"),
				ContainerPort:  jsii.Number(80),
			},
		},
		HealthCheckGracePeriodSeconds: jsii.Number(60),
	})
}