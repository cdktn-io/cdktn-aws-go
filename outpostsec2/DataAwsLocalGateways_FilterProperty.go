package outpostsec2


// Experimental.
type DataAwsLocalGateways_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateways#name DataAwsLocalGateways#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateways#values DataAwsLocalGateways#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

