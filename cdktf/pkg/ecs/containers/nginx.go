package containers

import (
	"fmt"
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/cloudwatchloggroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecsservice"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecstaskdefinition"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// This method will create the ECS task definition and the ECS service
func NewNginxContainer(stack *stack.MyStack) {
	logGroupName := fmt.Sprintf("ecs/%v-nginx", stack.Cfgs.AppName)
	cloudwatchloggroup.NewCloudwatchLogGroup(stack.TfStack, jsii.String(fmt.Sprintf("%v-log-group", logGroupName)), &cloudwatchloggroup.CloudwatchLogGroupConfig{
		Name:            jsii.String(logGroupName),
		RetentionInDays: jsii.Number(14),
	})

	containerName := fmt.Sprintf("%v-fna-nginx", stack.Cfgs.AppName)

	containerDefinitions := `
	[
		{
			"cpu": 256,
			"image": "nginx",
			"name": "<<CONTAINER_NAME>>",
			"portMappings": [{ "containerPort": 80 }],
			"logConfiguration": {
				"logDriver": "awslogs",
				"options": {
					"awslogs-group": "<<GROUP_NAME>>",
					"awslogs-region": "<<AWS_REGION>>",
					"awslogs-stream-prefix": "ecs"
				}
			}
		}
	]`
	containerDefinitions = strings.Replace(containerDefinitions, "<<CONTAINER_NAME>>", containerName, -1)
	containerDefinitions = strings.Replace(containerDefinitions, "<<GROUP_NAME>>", logGroupName, -1)
	containerDefinitions = strings.Replace(containerDefinitions, "<<AWS_REGION>>", stack.Cfgs.Region, -1)

	ecsNginxTaskDefinitionName := fmt.Sprintf("%v-ecs-nginx-td", stack.Cfgs.AppName)
	td := ecstaskdefinition.NewEcsTaskDefinition(stack.TfStack, jsii.String(ecsNginxTaskDefinitionName), &ecstaskdefinition.EcsTaskDefinitionConfig{
		Family:                  jsii.String(ecsNginxTaskDefinitionName),
		Cpu:                     jsii.String("256"),
		Memory:                  jsii.String("512"),
		NetworkMode:             jsii.String("awsvpc"),
		RequiresCompatibilities: jsii.Strings("FARGATE"),
		//
		// In this Execution role we need to add all the permissions to all services that
		// the container will need to execute
		// example we always need to have the cloud watch and secret manager role and others specific to each service
		ExecutionRoleArn:     stack.IAMCloudWatch.Role.Arn(),
		ContainerDefinitions: jsii.String(containerDefinitions),
	})

	ecsTaskDefinitionSecGroupName := fmt.Sprintf("%v-ecs-nginx-sec-group", stack.Cfgs.AppName)
	secGroup := securitygroup.NewSecurityGroup(stack.TfStack, jsii.String(ecsTaskDefinitionSecGroupName), &securitygroup.SecurityGroupConfig{
		VpcId: stack.Vpc.Id(),
		Ingress: &[]*securitygroup.SecurityGroupIngress{
			{
				Protocol:       jsii.String("TCP"),
				FromPort:       jsii.Number(0),
				ToPort:         jsii.Number(65535),
				SecurityGroups: &[]*string{stack.PublicAppLoadBalancer.SecGroup.Id()},
			},
		},
		Egress: &[]*securitygroup.SecurityGroupEgress{
			{
				Protocol:   jsii.String("TCP"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				FromPort:   jsii.Number(0),
				ToPort:     jsii.Number(65535),
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
			AssignPublicIp: jsii.Bool(true),
			Subnets:        &[]*string{stack.Subnets.PrivateA.Id(), stack.Subnets.PrivateB.Id()},
			SecurityGroups: &[]*string{secGroup.Id()},
		},
		LoadBalancer: &[]*ecsservice.EcsServiceLoadBalancer{
			{
				TargetGroupArn: stack.PublicAppLoadBalancer.TargetGroup.Arn(),
				ContainerName:  jsii.String(containerName),
				ContainerPort:  jsii.Number(80),
			},
		},
		HealthCheckGracePeriodSeconds: jsii.Number(60),
		ServiceConnectConfiguration: &ecsservice.EcsServiceServiceConnectConfiguration{
			Enabled:   jsii.Bool(true),
			Namespace: stack.ServiceDiscoveryPrivateNamespace.Arn(),
			Service: &[]*ecsservice.EcsServiceServiceConnectConfigurationService{
				{
					//ECS Service Name - The name used in the service not the name used in the container name
					DiscoveryName: jsii.String("nginx"),
					PortName:      jsii.String("80"),
				},
			},
		},
	})
}
