package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alb"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alblistener"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/albtargetgroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewPublicApplicationLoadBalancer(stack *stack.MyStack) {
	secGroupName := fmt.Sprintf("%v-alb-sec-group", stack.Cfgs.AppName)
	stack.PublicAppLoadBalancer.SecGroup = securitygroup.NewSecurityGroup(stack.TfStack, jsii.String(secGroupName), &securitygroup.SecurityGroupConfig{
		Description: jsii.String("Allows access from internet"),
		VpcId:       stack.Vpc.Id(),
		Ingress: []*securitygroup.SecurityGroupIngress{
			{
				Protocol:   jsii.String("TCP"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				FromPort:   jsii.Number(80),
				ToPort:     jsii.Number(80),
			},
			{
				Protocol:   jsii.String("TCP"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				FromPort:   jsii.Number(443),
				ToPort:     jsii.Number(443),
			},
		},
		Egress: []*securitygroup.SecurityGroupEgress{
			{
				Protocol:   jsii.String("TCP"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				FromPort:   jsii.Number(0),
				ToPort:     jsii.Number(65535),
			},
		},
	})

	albName := fmt.Sprintf("%v-alb", stack.Cfgs.AppName)
	stack.PublicAppLoadBalancer.Alb = alb.NewAlb(stack.TfStack, jsii.String(albName), &alb.AlbConfig{
		EnableHttp2:      true,
		Internal:         false,
		LoadBalancerType: jsii.String("application"),
		IpAddressType:    jsii.String("ipv4"),
		SubnetMapping: []*alb.AlbSubnetMapping{
			{
				SubnetId: stack.Subnets.PrivateA.Id(),
			},
			{
				SubnetId: stack.Subnets.Public.Id(),
			},
		},
		SecurityGroups: &[]*string{stack.PublicAppLoadBalancer.SecGroup.Id()},
	})

	albTargetGroupName := fmt.Sprintf("%v-alb-tg", stack.Cfgs.AppName)
	stack.PublicAppLoadBalancer.TargetGroup = albtargetgroup.NewAlbTargetGroup(stack.TfStack, jsii.String(albTargetGroupName), &albtargetgroup.AlbTargetGroupConfig{
		VpcId:           stack.Vpc.Id(),
		TargetType:      jsii.String("ip"),
		Protocol:        jsii.String("HTTP"),
		ProtocolVersion: jsii.String("HTTP1"),
		Port:            jsii.Number(80),
		HealthCheck: &albtargetgroup.AlbTargetGroupHealthCheck{
			Enabled: true,
			Path:    jsii.String("/"),
			Port:    jsii.String("80"),
		},
	})

	albListenerName := fmt.Sprintf("%v-alb-listener", stack.Cfgs.AppName)
	alblistener.NewAlbListener(stack.TfStack, jsii.String(albListenerName), &alblistener.AlbListenerConfig{
		LoadBalancerArn: stack.PublicAppLoadBalancer.Alb.Arn(),
		Protocol:        jsii.String("HTTP"),
		Port:            jsii.Number(80),
		DefaultAction: []*alblistener.AlbListenerDefaultAction{
			{
				Type:           jsii.String("forward"),
				TargetGroupArn: stack.PublicAppLoadBalancer.SecGroup.Arn(),
			},
		},
	})
}
