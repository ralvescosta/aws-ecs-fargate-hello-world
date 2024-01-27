package configs

type Configs struct {
	Provider *ProviderConfigs

	Vpc *VpcConfigs

	PrivateSubnetA *SubnetConfigs
	PrivateSubnetB *SubnetConfigs
	PublicSubnetA  *SubnetConfigs
	PublicSubnetB  *SubnetConfigs

	InternetGateway *InternetGatewayConfigs

	PrivateARouteTable *RouteTableConfigs
	PrivateBRouteTable *RouteTableConfigs
	PublicRouteTable   *RouteTableConfigs

	NatGatewayA *NatGatewayConfigs
	NatGatewayB *NatGatewayConfigs
}

type ProviderConfigs struct {
	AppId                    string
	Region                   string
	AccessKey                string
	SecretKey                string
	CloudBackendHostname     string
	CloudBackendOrganization string
}

type VpcConfigs struct {
	Name      string
	CidrBlock string
}

type SubnetConfigs struct {
	Name             string
	CidrBlock        string
	AvailabilityZone string
}

type InternetGatewayConfigs struct {
	Name string
}

type RouteTableConfigs struct {
	Name                   string
	SubnetAssociationNames []string
}

type NatGatewayConfigs struct {
	Name          string
	ElasticIpName string
}

func NewConfigs() *Configs {
	return &Configs{
		Provider: &ProviderConfigs{
			// Terraform cloud Project
			// REQUIRED
			AppId: "cdktf-hello-world",
			// AWS Region
			Region: "us-west-1",
			// AWS IAM Programmatic Access -  Access Key
			// REQUIRED
			AccessKey: "",
			// AWS IAM Programmatic Access -  Secret Key
			// REQUIRED
			SecretKey:            "",
			CloudBackendHostname: "app.terraform.io",
			// Terraform Cloud Organization
			// REQUIRED
			CloudBackendOrganization: "",
		},
		Vpc: &VpcConfigs{
			Name:      "fna-vpc",
			CidrBlock: "10.0.0.0/16",
		},
		PrivateSubnetA: &SubnetConfigs{
			Name:             "fna-private-a",
			CidrBlock:        "10.0.1.0/24",
			AvailabilityZone: "us-west-1a",
		},
		PrivateSubnetB: &SubnetConfigs{
			Name:             "fna-private-b",
			CidrBlock:        "10.0.2.0/24",
			AvailabilityZone: "us-west-1c",
		},
		PublicSubnetA: &SubnetConfigs{
			Name:             "fna-public-a",
			CidrBlock:        "10.0.3.0/24",
			AvailabilityZone: "us-west-1a",
		},
		PublicSubnetB: &SubnetConfigs{
			Name:             "fna-public-b",
			CidrBlock:        "10.0.4.0/24",
			AvailabilityZone: "us-west-1c",
		},
		InternetGateway: &InternetGatewayConfigs{
			Name: "fna-igw",
		},
		PrivateARouteTable: &RouteTableConfigs{
			Name:                   "fna-private-a-rt",
			SubnetAssociationNames: []string{"fna-private-rt-a"},
		},
		PrivateBRouteTable: &RouteTableConfigs{
			Name:                   "fna-private-b-rt",
			SubnetAssociationNames: []string{"fna-private-rt-b"},
		},
		PublicRouteTable: &RouteTableConfigs{
			Name:                   "fna-public-rt",
			SubnetAssociationNames: []string{"fna-public-rt-a", "fna-public-rt-b"},
		},
		NatGatewayA: &NatGatewayConfigs{
			Name:          "nt-gtw-a",
			ElasticIpName: "fna-eip-a",
		},
		NatGatewayB: &NatGatewayConfigs{
			Name:          "nt-gtw-b",
			ElasticIpName: "fna-eip-b",
		},
	}
}
