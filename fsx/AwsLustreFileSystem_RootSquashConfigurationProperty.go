package fsx


// Experimental.
type AwsLustreFileSystem_RootSquashConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#no_squash_nids AwsLustreFileSystem#no_squash_nids}.
	// Experimental.
	NoSquashNids *[]*string `field:"optional" json:"noSquashNids" yaml:"noSquashNids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#root_squash AwsLustreFileSystem#root_squash}.
	// Experimental.
	RootSquash *string `field:"optional" json:"rootSquash" yaml:"rootSquash"`
}

