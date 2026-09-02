package awsdirectoryservice


// Experimental.
type TfDirectory_ConnectSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#customer_dns_ips TfDirectory#customer_dns_ips}.
	// Experimental.
	CustomerDnsIps *[]*string `field:"required" json:"customerDnsIps" yaml:"customerDnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#customer_username TfDirectory#customer_username}.
	// Experimental.
	CustomerUsername *string `field:"required" json:"customerUsername" yaml:"customerUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#subnet_ids TfDirectory#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#vpc_id TfDirectory#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

