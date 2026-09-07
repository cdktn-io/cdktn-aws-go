package directoryservice


// Experimental.
type AwsDirectory_ConnectSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#customer_dns_ips AwsDirectory#customer_dns_ips}.
	// Experimental.
	CustomerDnsIps *[]*string `field:"required" json:"customerDnsIps" yaml:"customerDnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#customer_username AwsDirectory#customer_username}.
	// Experimental.
	CustomerUsername *string `field:"required" json:"customerUsername" yaml:"customerUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#subnet_ids AwsDirectory#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#vpc_id AwsDirectory#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

