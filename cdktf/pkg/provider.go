package pkg

import (
	"github.com/aws/jsii-runtime-go"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v18/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
)

func NewAWSScopeProvider(cfgs *configs.Configs) (cdktf.App, cdktf.TerraformStack) {
	appScope := cdktf.NewApp(nil)

	tfScope := cdktf.NewTerraformStack(appScope, jsii.String(cfgs.AppName))

	awsprovider.NewAwsProvider(tfScope, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region:    jsii.String(cfgs.Region),
		AccessKey: jsii.String(cfgs.AccessKey),
		SecretKey: jsii.String(cfgs.SecretKey),
	})

	cdktf.NewCloudBackend(tfScope, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String(cfgs.TerraformCloudHostname),
		Organization: jsii.String(cfgs.TerraformCloudOrganization),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String(cfgs.AppName)),
	})

	return appScope, tfScope
}
