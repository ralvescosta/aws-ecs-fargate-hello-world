package configs

import (
	"errors"
	"fmt"
	"os"
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

func NewConfigs() *Configs {
	return &Configs{
		AppName:  requiredEnv("APP_NAME"),
		LogLevel: envOrDefault("LOG_LEVEL", "debug"),

		Region:                     requiredEnv("AWS_REGION"),
		AccessKey:                  requiredEnv("AWS_ACCESS_KEY"),
		SecretKey:                  requiredEnv("AWS_SECRET_KEY"),
		TerraformCloudHostname:     requiredEnv("TERRAFORM_CLOUD_HOSTNAME"),
		TerraformCloudOrganization: requiredEnv("TERRAFORM_CLOUD_ORGANIZATION"),

		VpcCIDR:           requiredEnv("VPC_CIDR"),
		PrivateSubnetCIDR: requiredEnv("PRIVATE_SUBNET_CIDR"),
		PrivateSubnetAZ:   requiredEnv("PRIVATE_SUBNET_AZ"),
		PublicSubnetCIDR:  requiredEnv("PUBLIC_SUBNET_CIDR"),
		PublicSubnetAZ:    requiredEnv("PUBLIC_SUBNET_AZ"),
	}
}

func requiredEnv(envKey string) string {
	value := os.Getenv(envKey)

	if value == "" {
		panic(errors.New(fmt.Sprintf("env %v is required, but was founded empty", envKey)))
	}

	return value
}

func envOrDefault(envKey, def string) string {
	value := os.Getenv(envKey)

	if value != "" {
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
