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

type MyStack struct {
	Cfgs   *configs.Configs
	Logger *zap.SugaredLogger

	TfStack cdktf.TerraformStack

	Vpc                            vpc.Vpc
	PrivateSubnet                  subnet.Subnet
	PublicSubnet                   subnet.Subnet
	InternetGateway                internetgateway.InternetGateway
	NatGateway                     natgateway.NatGateway
	PrivateRouteTable              routetable.RouteTable
	PublicRouteTable               routetable.RouteTable
	ElasticLoadBalancerSecGroup    securitygroup.SecurityGroup
	ElasticLoadBalancerTargetGroup albtargetgroup.AlbTargetGroup
	ElasticLoadBalancer            alb.Alb
	EcsCluster                     ecscluster.EcsCluster
	EcrAPIRepository               ecrrepository.EcrRepository
	EcrGrpcRepository              ecrrepository.EcrRepository
}
