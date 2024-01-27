package network

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/internetgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/natgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetableassociation"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func NewRouteTables(
	cfgs *configs.Configs,
	tfStack cdktf.TerraformStack,
	fnaVpc vpc.Vpc,
	igw internetgateway.InternetGateway,
	privateA subnet.Subnet,
	privateB subnet.Subnet,
	publicA subnet.Subnet,
	publicB subnet.Subnet,
	natGtwA natgateway.NatGateway,
	natGtwB natgateway.NatGateway,
) (
	privateARouteTable routetable.RouteTable,
	privateBRouteTable routetable.RouteTable,
	publicRouteTable routetable.RouteTable,
) {
	privateARouteTable = routetable.NewRouteTable(tfStack, jsii.String(cfgs.PrivateARouteTable.Name), &routetable.RouteTableConfig{
		VpcId: fnaVpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: natGtwA.Id(),
			},
		},
	})

	privateBRouteTable = routetable.NewRouteTable(tfStack, jsii.String(cfgs.PrivateBRouteTable.Name), &routetable.RouteTableConfig{
		VpcId: fnaVpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: natGtwB.Id(),
			},
		},
	})

	publicRouteTable = routetable.NewRouteTable(tfStack, jsii.String(cfgs.PublicRouteTable.Name), &routetable.RouteTableConfig{
		VpcId: fnaVpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: igw.Id(),
			},
		},
	})

	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PrivateARouteTable.SubnetAssociationNames[0]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: privateARouteTable.Id(),
		SubnetId:     privateA.Id(),
	})
	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PrivateBRouteTable.SubnetAssociationNames[0]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: privateBRouteTable.Id(),
		SubnetId:     privateB.Id(),
	})

	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PublicRouteTable.SubnetAssociationNames[0]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: publicRouteTable.Id(),
		SubnetId:     publicA.Id(),
	})
	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PublicRouteTable.SubnetAssociationNames[1]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: publicRouteTable.Id(),
		SubnetId:     publicB.Id(),
	})

	return
}
