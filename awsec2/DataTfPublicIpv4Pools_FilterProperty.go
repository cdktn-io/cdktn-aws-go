package awsec2


// Experimental.
type DataTfPublicIpv4Pools_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_public_ipv4_pools#name DataTfPublicIpv4Pools#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_public_ipv4_pools#values DataTfPublicIpv4Pools#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

