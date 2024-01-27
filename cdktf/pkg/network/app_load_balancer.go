package network

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alb"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alblistener"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/albtargetgroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func NewApplicationLoadBalancer(
	cfg *configs.Configs,
	tfStack cdktf.TerraformStack,
	fnaVpc vpc.Vpc,
	publicA subnet.Subnet,
	publicB subnet.Subnet,
) (albSecurityGroup securitygroup.SecurityGroup, targetGroup albtargetgroup.AlbTargetGroup, appLoadBalancer alb.Alb) {
	albSecurityGroup = securitygroup.NewSecurityGroup(tfStack, jsii.String("fna-alb-sg"), &securitygroup.SecurityGroupConfig{
		Description: jsii.String("Allows access from internet"),
		VpcId:       fnaVpc.Id(),
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

	appLoadBalancer = alb.NewAlb(tfStack, jsii.String("fna-alb"), &alb.AlbConfig{
		EnableHttp2:      true,
		Internal:         false,
		LoadBalancerType: jsii.String("application"),
		IpAddressType:    jsii.String("ipv4"),
		SubnetMapping: []*alb.AlbSubnetMapping{
			{
				SubnetId: publicA.Id(),
			},
			{
				SubnetId: publicB.Id(),
			},
		},
		SecurityGroups: &[]*string{albSecurityGroup.Id()},
	})

	targetGroup = albtargetgroup.NewAlbTargetGroup(tfStack, jsii.String("fna-alb-tg"), &albtargetgroup.AlbTargetGroupConfig{
		VpcId:           fnaVpc.Id(),
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

	alblistener.NewAlbListener(tfStack, jsii.String("fna-alb-l"), &alblistener.AlbListenerConfig{
		LoadBalancerArn: appLoadBalancer.Arn(),
		Protocol:        jsii.String("HTTP"),
		Port:            jsii.Number(80),
		DefaultAction: []*alblistener.AlbListenerDefaultAction{
			{
				Type:           jsii.String("forward"),
				TargetGroupArn: targetGroup.Arn(),
			},
		},
	})

	return
}
