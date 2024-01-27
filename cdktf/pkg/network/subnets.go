package network

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func NewSubnets(cfgs *configs.Configs, tfStack cdktf.TerraformStack, fnaVpc vpc.Vpc) (
	privateA subnet.Subnet, privateB subnet.Subnet, publicA subnet.Subnet, publicB subnet.Subnet,
) {
	privateA = subnet.NewSubnet(tfStack, jsii.Sprintf(cfgs.PrivateSubnetA.Name), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.String(cfgs.PrivateSubnetA.CidrBlock),
		AvailabilityZone: jsii.String(cfgs.PrivateSubnetA.AvailabilityZone),
	})

	privateB = subnet.NewSubnet(tfStack, jsii.Sprintf(cfgs.PrivateSubnetB.Name), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.Sprintf(cfgs.PrivateSubnetB.CidrBlock),
		AvailabilityZone: jsii.String(cfgs.PrivateSubnetB.AvailabilityZone),
	})

	publicA = subnet.NewSubnet(tfStack, jsii.Sprintf(cfgs.PublicSubnetA.Name), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.Sprintf(cfgs.PublicSubnetA.CidrBlock),
		AvailabilityZone: jsii.String(cfgs.PublicSubnetA.AvailabilityZone),
	})

	publicB = subnet.NewSubnet(tfStack, jsii.Sprintf(cfgs.PublicSubnetB.Name), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.Sprintf(cfgs.PublicSubnetB.CidrBlock),
		AvailabilityZone: jsii.String(cfgs.PublicSubnetB.AvailabilityZone),
	})

	return
}
