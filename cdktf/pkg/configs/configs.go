package configs

import (
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
)

type Configs struct {
	AppName  string
	LogLevel string

	Region                     string
	AccessKey                  string
	SecretKey                  string
	TerraformCloudHostname     string
	TerraformCloudOrganization string

	VpcCIDR string

	PrivateSubnetCIDR string
	PrivateSubnetAZ   string

	PublicSubnetCIDR string
	PublicSubnetAZ   string
}

func NewConfigs(logger *zap.SugaredLogger) *Configs {
	return &Configs{
		AppName:  requiredEnv(logger, "APP_NAME"),
		LogLevel: envOrDefault(logger, "LOG_LEVEL", "debug"),

		Region:                     requiredEnv(logger, "AWS_REGION"),
		AccessKey:                  requiredEnv(logger, "AWS_ACCESS_KEY"),
		SecretKey:                  requiredEnv(logger, "AWS_SECRET_KEY"),
		TerraformCloudHostname:     requiredEnv(logger, "TERRAFORM_CLOUD_HOSTNAME"),
		TerraformCloudOrganization: requiredEnv(logger, "TERRAFORM_CLOUD_ORGANIZATION"),

		VpcCIDR:           requiredEnv(logger, "VPC_CIDR"),
		PrivateSubnetCIDR: requiredEnv(logger, "PRIVATE_SUBNET_CIDR"),
		PrivateSubnetAZ:   requiredEnv(logger, "PRIVATE_SUBNET_AZ"),
		PublicSubnetCIDR:  requiredEnv(logger, "PUBLIC_SUBNET_CIDR"),
		PublicSubnetAZ:    requiredEnv(logger, "PUBLIC_SUBNET_AZ"),
	}
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
