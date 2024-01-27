package ecs

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alb"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/albtargetgroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscluster"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecsservice"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecstaskdefinition"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func NewECSFargate(
	cfgs *configs.Configs,
	tfStack cdktf.TerraformStack,
	fnaVpc vpc.Vpc,
	privateA subnet.Subnet,
	privateB subnet.Subnet,
	fnaAlbTargetGroup albtargetgroup.AlbTargetGroup,
	fnaAlbSecGroup securitygroup.SecurityGroup,
	fnaAlb alb.Alb,
) {
	cluster := ecscluster.NewEcsCluster(tfStack, jsii.String("fna-ecs-cluster"), &ecscluster.EcsClusterConfig{
		Name: jsii.String("fna-ecs-cluster"),
		Setting: []*ecscluster.EcsClusterSetting{
			{
				Name:  jsii.String("containerInsights"),
				Value: jsii.String("enabled"),
			},
		},
	})

	td := ecstaskdefinition.NewEcsTaskDefinition(tfStack, jsii.String("fna-td"), &ecstaskdefinition.EcsTaskDefinitionConfig{
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

	secGroup := securitygroup.NewSecurityGroup(tfStack, jsii.String("fna-ecs-sg"), &securitygroup.SecurityGroupConfig{
		Ingress: &[]*securitygroup.SecurityGroupIngress{
			{
				Protocol:       jsii.String("tcp"),
				FromPort:       jsii.Number(0),
				ToPort:         jsii.Number(6553),
				SecurityGroups: &[]*string{fnaAlbSecGroup.Id()},
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

	ecsservice.NewEcsService(tfStack, jsii.String("fna-svc"), &ecsservice.EcsServiceConfig{
		Name:           jsii.String("fna-svc"),
		Cluster:        cluster.Id(),
		TaskDefinition: td.Arn(),
		LaunchType:     jsii.String("FARGATE"),
		DesiredCount:   jsii.Number(2),
		DeploymentController: &ecsservice.EcsServiceDeploymentController{
			Type: jsii.String("ECS"),
		},
		NetworkConfiguration: &ecsservice.EcsServiceNetworkConfiguration{
			Subnets:        &[]*string{privateA.Id(), privateB.Id()},
			SecurityGroups: &[]*string{secGroup.Id()},
		},
		LoadBalancer: &[]*ecsservice.EcsServiceLoadBalancer{
			{
				ElbName:        fnaAlb.Name(),
				TargetGroupArn: fnaAlbTargetGroup.Arn(),
				ContainerName:  jsii.String("fna-nginx"),
				ContainerPort:  jsii.Number(80),
			},
		},
		HealthCheckGracePeriodSeconds: jsii.Number(60),
	})
}
