package awsvpclattice


// Experimental.
type AwsVpclatticeResourceConfiguration_DnsResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#domain_name AwsVpclatticeResourceConfiguration#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#ip_address_type AwsVpclatticeResourceConfiguration#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"required" json:"ipAddressType" yaml:"ipAddressType"`
}

