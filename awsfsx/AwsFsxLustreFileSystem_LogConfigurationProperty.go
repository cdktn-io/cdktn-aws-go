package awsfsx


// Experimental.
type AwsFsxLustreFileSystem_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#destination AwsFsxLustreFileSystem#destination}.
	// Experimental.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#level AwsFsxLustreFileSystem#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

