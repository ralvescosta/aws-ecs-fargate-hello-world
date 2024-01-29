package ecs

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscluster"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecsservice"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecstaskdefinition"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewECSFargate(stack *stack.MyStack) {
	cluster := ecscluster.NewEcsCluster(stack.TfStack, jsii.String("fna-ecs-cluster"), &ecscluster.EcsClusterConfig{
		Name: jsii.String("fna-ecs-cluster"),
		Setting: []*ecscluster.EcsClusterSetting{
			{
				Name:  jsii.String("containerInsights"),
				Value: jsii.String("enabled"),
			},
		},
	})

	td := ecstaskdefinition.NewEcsTaskDefinition(stack.TfStack, jsii.String("fna-td"), &ecstaskdefinition.EcsTaskDefinitionConfig{
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

	secGroup := securitygroup.NewSecurityGroup(stack.TfStack, jsii.String("fna-ecs-sg"), &securitygroup.SecurityGroupConfig{
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

	ecsservice.NewEcsService(stack.TfStack, jsii.String("fna-svc"), &ecsservice.EcsServiceConfig{
		Name:           jsii.String("fna-svc"),
		Cluster:        cluster.Id(),
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
