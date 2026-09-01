package awsebs


// Experimental.
type DataAwsEbsVolume_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_volume#name DataAwsEbsVolume#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_volume#values DataAwsEbsVolume#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

