package network

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/internetgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func NewInternetGateway(cfgs *configs.Configs, tfStack cdktf.TerraformStack, fnaVpc vpc.Vpc) (igw internetgateway.InternetGateway) {
	igw = internetgateway.NewInternetGateway(tfStack, jsii.String(cfgs.InternetGateway.Name), &internetgateway.InternetGatewayConfig{
		VpcId: fnaVpc.Id(),
	})

	return
}
