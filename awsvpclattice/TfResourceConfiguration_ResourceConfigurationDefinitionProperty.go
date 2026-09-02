package awsvpclattice


// Experimental.
type TfResourceConfiguration_ResourceConfigurationDefinitionProperty struct {
	// arn_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#arn_resource TfResourceConfiguration#arn_resource}
	// Experimental.
	ArnResource interface{} `field:"optional" json:"arnResource" yaml:"arnResource"`
	// dns_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#dns_resource TfResourceConfiguration#dns_resource}
	// Experimental.
	DnsResource interface{} `field:"optional" json:"dnsResource" yaml:"dnsResource"`
	// ip_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#ip_resource TfResourceConfiguration#ip_resource}
	// Experimental.
	IpResource interface{} `field:"optional" json:"ipResource" yaml:"ipResource"`
}

