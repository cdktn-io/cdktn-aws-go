package awsfsx


// Experimental.
type TfOpenzfsFileSystem_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#create TfOpenzfsFileSystem#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#delete TfOpenzfsFileSystem#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#update TfOpenzfsFileSystem#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

