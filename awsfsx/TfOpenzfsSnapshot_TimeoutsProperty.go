package awsfsx


// Experimental.
type TfOpenzfsSnapshot_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_snapshot#create TfOpenzfsSnapshot#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_snapshot#delete TfOpenzfsSnapshot#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_snapshot#read TfOpenzfsSnapshot#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_snapshot#update TfOpenzfsSnapshot#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

