package fsx


// Experimental.
type DataAwsOpenzfsSnapshot_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/fsx_openzfs_snapshot#name DataAwsOpenzfsSnapshot#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/fsx_openzfs_snapshot#values DataAwsOpenzfsSnapshot#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

