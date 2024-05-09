package stack

import (
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/alb"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/albtargetgroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecrrepository"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscluster"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/eip"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/iampolicy"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/iamrole"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/iamrolepolicyattachment"
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
		PublicA  subnet.Subnet
		PublicB  subnet.Subnet
	}

	RouteTable struct {
		PrivateA routetable.RouteTable
		PrivateB routetable.RouteTable
		PublicA  routetable.RouteTable
		PublicB  routetable.RouteTable
	}

	NatGateway struct {
		EIpA     eip.Eip
		PrivateA natgateway.NatGateway
		EIpB     eip.Eip
		PrivateB natgateway.NatGateway
	}

	IAMCloudWatch struct {
		Role          iamrole.IamRole
		Policy        iampolicy.IamPolicy
		RolePolicyAtt iamrolepolicyattachment.IamRolePolicyAttachment
	}

	CloudWatchLogs struct {
	}

	MyStack struct {
		Cfgs   *configs.Configs
		Logger *zap.SugaredLogger

		TfStack cdktf.TerraformStack

		Vpc                    vpc.Vpc
		Subnets                *Subnet
		InternetGateway        internetgateway.InternetGateway
		NatGateways            *NatGateway
		RouteTables            *RouteTable
		PublicAppLoadBalancer  *ApplicationLoadBalancer
		PrivateAppLoadBalancer *ApplicationLoadBalancer
		IAMCloudWatch          *IAMCloudWatch
		EcsCluster             ecscluster.EcsCluster
		EcrAPIRepository       ecrrepository.EcrRepository
		EcrGrpcRepository      ecrrepository.EcrRepository
	}
)
