package network

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetableassociation"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/stack"
)

func NewRouteTables(stack *stack.MyStack) {
	privateRouteTableName := fmt.Sprintf("%v-private-rt", stack.Cfgs.AppName)
	stack.PrivateRouteTable = routetable.NewRouteTable(stack.TfStack, jsii.String(privateRouteTableName), &routetable.RouteTableConfig{
		VpcId: stack.Vpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: stack.NatGateway.Id(),
			},
		},
	})

	publicRouteTableName := fmt.Sprintf("%v-public-rt", stack.Cfgs.AppName)
	stack.PublicRouteTable = routetable.NewRouteTable(stack.TfStack, jsii.String(publicRouteTableName), &routetable.RouteTableConfig{
		VpcId: stack.Vpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: stack.InternetGateway.Id(),
			},
		},
	})

	privateRouteTableAssociationName := fmt.Sprintf("%v-private-rt-a", stack.Cfgs.AppName)
	routetableassociation.NewRouteTableAssociation(stack.TfStack, jsii.String(privateRouteTableAssociationName), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: stack.PrivateRouteTable.Id(),
		SubnetId:     stack.Subnets.PrivateA.Id(),
	})

	publicRouteTableAssociationName := fmt.Sprintf("%v-public-rt-a", stack.Cfgs.AppName)
	routetableassociation.NewRouteTableAssociation(stack.TfStack, jsii.String(publicRouteTableAssociationName), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: stack.PublicRouteTable.Id(),
		SubnetId:     stack.Subnets.Public.Id(),
	})
}
