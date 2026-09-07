package vpclattice


// Experimental.
type AwsResourceConfiguration_ResourceConfigurationDefinitionProperty struct {
	// arn_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#arn_resource AwsResourceConfiguration#arn_resource}
	// Experimental.
	ArnResource interface{} `field:"optional" json:"arnResource" yaml:"arnResource"`
	// dns_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#dns_resource AwsResourceConfiguration#dns_resource}
	// Experimental.
	DnsResource interface{} `field:"optional" json:"dnsResource" yaml:"dnsResource"`
	// ip_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#ip_resource AwsResourceConfiguration#ip_resource}
	// Experimental.
	IpResource interface{} `field:"optional" json:"ipResource" yaml:"ipResource"`
}

