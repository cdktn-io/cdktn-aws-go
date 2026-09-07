package fsx


// Experimental.
type AwsOpenzfsFileSystem_ClientConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#clients AwsOpenzfsFileSystem#clients}.
	// Experimental.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#options AwsOpenzfsFileSystem#options}.
	// Experimental.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

