package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetableassociation"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

// This method will create the route tables for each subnet and associate the route tables with than
//
// The route table bound the internet gateway to the public subnets and the nat gateway to the private subnets
func NewRouteTables(stack *stack.MyStack) {
	privateRouteTableAName := fmt.Sprintf("%v-private-rt-a", stack.Cfgs.AppName)
	stack.RouteTables.PrivateA = routetable.NewRouteTable(stack.TfStack, jsii.String(privateRouteTableAName), &routetable.RouteTableConfig{
		VpcId: stack.Vpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: stack.NatGateways.PrivateA.Id(),
			},
		},
	})

	privateRouteTableBName := fmt.Sprintf("%v-private-rt-b", stack.Cfgs.AppName)
	stack.RouteTables.PrivateB = routetable.NewRouteTable(stack.TfStack, jsii.String(privateRouteTableBName), &routetable.RouteTableConfig{
		VpcId: stack.Vpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: stack.NatGateways.PrivateB.Id(),
			},
		},
	})

	publicRouteTableAName := fmt.Sprintf("%v-public-rt-a", stack.Cfgs.AppName)
	stack.RouteTables.PublicA = routetable.NewRouteTable(stack.TfStack, jsii.String(publicRouteTableAName), &routetable.RouteTableConfig{
		VpcId: stack.Vpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: stack.InternetGateway.Id(),
			},
		},
	})

	publicRouteTableBName := fmt.Sprintf("%v-public-rt-b", stack.Cfgs.AppName)
	stack.RouteTables.PublicB = routetable.NewRouteTable(stack.TfStack, jsii.String(publicRouteTableBName), &routetable.RouteTableConfig{
		VpcId: stack.Vpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: stack.InternetGateway.Id(),
			},
		},
	})

	privateRouteTableAssociationAName := fmt.Sprintf("%v-private-a-rta", stack.Cfgs.AppName)
	routetableassociation.NewRouteTableAssociation(stack.TfStack, jsii.String(privateRouteTableAssociationAName), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: stack.RouteTables.PrivateA.Id(),
		SubnetId:     stack.Subnets.PrivateA.Id(),
	})

	privateRouteTableAssociationBName := fmt.Sprintf("%v-private-b-rta", stack.Cfgs.AppName)
	routetableassociation.NewRouteTableAssociation(stack.TfStack, jsii.String(privateRouteTableAssociationBName), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: stack.RouteTables.PrivateB.Id(),
		SubnetId:     stack.Subnets.PrivateB.Id(),
	})

	publicRouteTableAssociationAName := fmt.Sprintf("%v-public-a-rta", stack.Cfgs.AppName)
	routetableassociation.NewRouteTableAssociation(stack.TfStack, jsii.String(publicRouteTableAssociationAName), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: stack.RouteTables.PublicA.Id(),
		SubnetId:     stack.Subnets.PublicA.Id(),
	})

	publicRouteTableAssociationBName := fmt.Sprintf("%v-public-b-rta", stack.Cfgs.AppName)
	routetableassociation.NewRouteTableAssociation(stack.TfStack, jsii.String(publicRouteTableAssociationBName), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: stack.RouteTables.PublicB.Id(),
		SubnetId:     stack.Subnets.PublicB.Id(),
	})
}
