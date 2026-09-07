package fsx


// Experimental.
type AwsOpenzfsFileSystem_ReadCacheConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#size AwsOpenzfsFileSystem#size}.
	// Experimental.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#sizing_mode AwsOpenzfsFileSystem#sizing_mode}.
	// Experimental.
	SizingMode *string `field:"optional" json:"sizingMode" yaml:"sizingMode"`
}

