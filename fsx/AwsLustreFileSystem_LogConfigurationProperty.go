package fsx


// Experimental.
type AwsLustreFileSystem_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#destination AwsLustreFileSystem#destination}.
	// Experimental.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#level AwsLustreFileSystem#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

