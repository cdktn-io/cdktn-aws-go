package awsfsx


// Experimental.
type AwsFsxLustreFileSystem_DataReadCacheConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#sizing_mode AwsFsxLustreFileSystem#sizing_mode}.
	// Experimental.
	SizingMode *string `field:"required" json:"sizingMode" yaml:"sizingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#size AwsFsxLustreFileSystem#size}.
	// Experimental.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

