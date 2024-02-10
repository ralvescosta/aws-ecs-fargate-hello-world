package stack

import (
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alb"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/albtargetgroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecrrepository"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscluster"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/internetgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/natgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/securitygroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
	"go.uber.org/zap"
)

type (
	ApplicationLoadBalancer struct {
		SecGroup    securitygroup.SecurityGroup
		TargetGroup albtargetgroup.AlbTargetGroup
		Alb         alb.Alb
	}

	Subnet struct {
		PrivateA subnet.Subnet
		PrivateB subnet.Subnet
		Public   subnet.Subnet
	}

	MyStack struct {
		Cfgs   *configs.Configs
		Logger *zap.SugaredLogger

		TfStack cdktf.TerraformStack

		Vpc                    vpc.Vpc
		Subnets                *Subnet
		InternetGateway        internetgateway.InternetGateway
		NatGateway             natgateway.NatGateway
		PrivateRouteTable      routetable.RouteTable
		PublicRouteTable       routetable.RouteTable
		PublicAppLoadBalancer  *ApplicationLoadBalancer
		PrivateAppLoadBalancer *ApplicationLoadBalancer
		EcsCluster             ecscluster.EcsCluster
		EcrAPIRepository       ecrrepository.EcrRepository
		EcrGrpcRepository      ecrrepository.EcrRepository
	}
)
