package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/ralvescosta/dotenv"
	"go.uber.org/zap"
)

type Configs struct {
	AppName  string
	LogLevel string

	Region                     string
	TerraformCloudHostname     string
	TerraformCloudOrganization string

	VpcCIDR string

	PrivateSubnetA_CIDR string
	PrivateSubnetA_AZ   string
	PrivateSubnetB_CIDR string
	PrivateSubnetB_AZ   string

	PublicSubnetA_CIDR string
	PublicSubnetA_AZ   string
	PublicSubnetB_CIDR string
	PublicSubnetB_AZ   string
}

func NewConfigs() (*Configs, *zap.SugaredLogger) {
	logger := zap.S().Named("cdktf")

	goEnv := envOrDefault(logger, "GO_ENV", "staging")

	dotenv.Configure(fmt.Sprintf("./.env.%v", goEnv))

	logLevel := envOrDefault(logger, "LOG_LEVEL", "debug")
	appName := requiredEnv(logger, "APP_NAME")

	return &Configs{
		AppName:  appName,
		LogLevel: logLevel,

		Region:                     requiredEnv(logger, "AWS_REGION"),
		TerraformCloudHostname:     requiredEnv(logger, "TERRAFORM_CLOUD_HOSTNAME"),
		TerraformCloudOrganization: requiredEnv(logger, "TERRAFORM_CLOUD_ORGANIZATION"),

		VpcCIDR:             requiredEnv(logger, "VPC_CIDR"),
		PrivateSubnetA_CIDR: requiredEnv(logger, "PRIVATE_SUBNET_A_CIDR"),
		PrivateSubnetA_AZ:   requiredEnv(logger, "PRIVATE_SUBNET_A_AZ"),
		PrivateSubnetB_CIDR: requiredEnv(logger, "PRIVATE_SUBNET_B_CIDR"),
		PrivateSubnetB_AZ:   requiredEnv(logger, "PRIVATE_SUBNET_B_AZ"),
		PublicSubnetA_CIDR:  requiredEnv(logger, "PUBLIC_SUBNET_A_CIDR"),
		PublicSubnetA_AZ:    requiredEnv(logger, "PUBLIC_SUBNET_A_AZ"),
		PublicSubnetB_CIDR:  requiredEnv(logger, "PUBLIC_SUBNET_B_CIDR"),
		PublicSubnetB_AZ:    requiredEnv(logger, "PUBLIC_SUBNET_B_AZ"),
	}, logger
}

func requiredEnv(logger *zap.SugaredLogger, envKey string) string {
	value := os.Getenv(envKey)

	if value == "" {
		logger.Panic(errors.New(fmt.Sprintf("env %v is required, but was founded empty", envKey)))
	}

	return value
}

func envOrDefault(logger *zap.SugaredLogger, envKey, def string) string {
	value := os.Getenv(envKey)

	if value != "" {
		logger.Debug(fmt.Sprintf("env key %v without value, assuming the default value", envKey))
		return value
	}

	return def
}

// func NewConfigs() *Configs {
// 	return &Configs{
// 		Provider: &ProviderConfigs{
// 			// Terraform cloud Project
// 			// REQUIRED
// 			AppId: "cdktf-hello-world",
// 			// AWS Region
// 			Region: "us-west-1",
// 			// AWS IAM Programmatic Access -  Access Key
// 			// REQUIRED
// 			AccessKey: "",
// 			// AWS IAM Programmatic Access -  Secret Key
// 			// REQUIRED
// 			SecretKey:            "",
// 			CloudBackendHostname: "app.terraform.io",
// 			// Terraform Cloud Organization
// 			// REQUIRED
// 			CloudBackendOrganization: "",
// 		},
// 		Vpc: &VpcConfigs{
// 			Name:      "fna-vpc",
// 			CidrBlock: "10.0.0.0/16",
// 		},
// 		PrivateSubnetA: &SubnetConfigs{
// 			Name:             "fna-private-a",
// 			CidrBlock:        "10.0.1.0/24",
// 			AvailabilityZone: "us-west-1a",
// 		},
// 		PrivateSubnetB: &SubnetConfigs{
// 			Name:             "fna-private-b",
// 			CidrBlock:        "10.0.2.0/24",
// 			AvailabilityZone: "us-west-1c",
// 		},
// 		PublicSubnetA: &SubnetConfigs{
// 			Name:             "fna-public-a",
// 			CidrBlock:        "10.0.3.0/24",
// 			AvailabilityZone: "us-west-1a",
// 		},
// 		PublicSubnetB: &SubnetConfigs{
// 			Name:             "fna-public-b",
// 			CidrBlock:        "10.0.4.0/24",
// 			AvailabilityZone: "us-west-1c",
// 		},
// 		InternetGateway: &InternetGatewayConfigs{
// 			Name: "fna-igw",
// 		},
// 		PrivateARouteTable: &RouteTableConfigs{
// 			Name:                   "fna-private-a-rt",
// 			SubnetAssociationNames: []string{"fna-private-rt-a"},
// 		},
// 		PrivateBRouteTable: &RouteTableConfigs{
// 			Name:                   "fna-private-b-rt",
// 			SubnetAssociationNames: []string{"fna-private-rt-b"},
// 		},
// 		PublicRouteTable: &RouteTableConfigs{
// 			Name:                   "fna-public-rt",
// 			SubnetAssociationNames: []string{"fna-public-rt-a", "fna-public-rt-b"},
// 		},
// 		NatGatewayA: &NatGatewayConfigs{
// 			Name:          "nt-gtw-a",
// 			ElasticIpName: "fna-eip-a",
// 		},
// 		NatGatewayB: &NatGatewayConfigs{
// 			Name:          "nt-gtw-b",
// 			ElasticIpName: "fna-eip-b",
// 		},
// 	}
// }
