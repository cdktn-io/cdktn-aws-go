package ebs


// Experimental.
type DataAwsVolume_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_volume#name DataAwsVolume#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_volume#values DataAwsVolume#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

