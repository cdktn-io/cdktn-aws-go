package transitgateway


// Experimental.
type DataAwsMulticastDomain_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_multicast_domain#name DataAwsMulticastDomain#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_multicast_domain#values DataAwsMulticastDomain#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

