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

func NewApplicationLoadBalancer(stack *stack.MyStack) {
	secGroupName := fmt.Sprintf("%v-sec-group", stack.Cfgs.AppName)
	stack.ElasticLoadBalancerSecGroup = securitygroup.NewSecurityGroup(stack.TfStack, jsii.String(secGroupName), &securitygroup.SecurityGroupConfig{
		Description: jsii.String("Allows access from internet"),
		VpcId:       stack.Vpc.Id(),
		Ingress: []*securitygroup.SecurityGroupIngress{
			{
				Protocol:   jsii.String("HTTP"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				ToPort:     jsii.Number(80),
				FromPort:   jsii.Number(80),
			},
		},
		Egress: []*securitygroup.SecurityGroupEgress{
			{
				Protocol:   jsii.String("HTTP"),
				CidrBlocks: jsii.Strings("0.0.0.0/0"),
				ToPort:     jsii.Number(0),
				FromPort:   jsii.Number(65_535),
			},
		},
	})

	albName := fmt.Sprintf("%v-alb", stack.Cfgs.AppName)
	stack.ElasticLoadBalancer = alb.NewAlb(stack.TfStack, jsii.String(albName), &alb.AlbConfig{
		EnableHttp2:      true,
		Internal:         false,
		LoadBalancerType: jsii.String("application"),
		IpAddressType:    jsii.String("ipv4"),
		SubnetMapping: []*alb.AlbSubnetMapping{
			{
				SubnetId: stack.PublicSubnet.Id(),
			},
		},
		SecurityGroups: &[]*string{stack.ElasticLoadBalancerSecGroup.Id()},
	})

	albTargetGroupName := fmt.Sprintf("%v-alb-tg", stack.Cfgs.AppName)
	stack.ElasticLoadBalancerTargetGroup = albtargetgroup.NewAlbTargetGroup(stack.TfStack, jsii.String(albTargetGroupName), &albtargetgroup.AlbTargetGroupConfig{
		VpcId:           stack.Vpc.Id(),
		TargetType:      jsii.String("ip"),
		Protocol:        jsii.String("HTTP"),
		ProtocolVersion: jsii.String("HTTP1"),
		Port:            jsii.Number(80),
		HealthCheck: &albtargetgroup.AlbTargetGroupHealthCheck{
			Enabled: true,
			Path:    jsii.String("/health"),
			Port:    jsii.String("80"),
		},
	})

	albListenerName := fmt.Sprintf("%v-alb-listener", stack.Cfgs.AppName)
	alblistener.NewAlbListener(stack.TfStack, jsii.String(albListenerName), &alblistener.AlbListenerConfig{
		LoadBalancerArn: stack.ElasticLoadBalancer.Arn(),
		Protocol:        jsii.String("HTTP"),
		Port:            jsii.Number(80),
		DefaultAction: []*alblistener.AlbListenerDefaultAction{
			{
				Type:           jsii.String("forward"),
				TargetGroupArn: stack.ElasticLoadBalancerSecGroup.Arn(),
			},
		},
	})
}
