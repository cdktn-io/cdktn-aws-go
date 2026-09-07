package outpostsec2


// Experimental.
type DataAwsCoipPools_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_coip_pools#name DataAwsCoipPools#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_coip_pools#values DataAwsCoipPools#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

