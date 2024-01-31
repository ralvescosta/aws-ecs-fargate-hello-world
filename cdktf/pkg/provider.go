package pkg

import (
	"context"

	awsCredentialsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/jsii-runtime-go"
	awsProvider "github.com/cdktf/cdktf-provider-aws-go/aws/v18/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/cdktf/pkg/configs"
	"go.uber.org/zap"
)

func NewAWSScopeProvider(logger *zap.SugaredLogger, cfgs *configs.Configs) (cdktf.App, cdktf.TerraformStack) {
	appScope := cdktf.NewApp(nil)
	tfScope := cdktf.NewTerraformStack(appScope, jsii.String(cfgs.AppName))

	ctx := context.Background()
	awsConfigs, err := awsCredentialsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	awsCredentials, err := awsConfigs.Credentials.Retrieve(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	awsProvider.NewAwsProvider(tfScope, jsii.String("AWS"), &awsProvider.AwsProviderConfig{
		Region:    jsii.String(cfgs.Region),
		AccessKey: jsii.String(awsCredentials.AccessKeyID),
		SecretKey: jsii.String(awsCredentials.SecretAccessKey),
	})

	cdktf.NewCloudBackend(tfScope, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String(cfgs.TerraformCloudHostname),
		Organization: jsii.String(cfgs.TerraformCloudOrganization),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String(cfgs.AppName)),
	})

	return appScope, tfScope
}
