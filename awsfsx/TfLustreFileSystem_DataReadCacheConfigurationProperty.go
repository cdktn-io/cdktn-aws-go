package awsfsx


// Experimental.
type TfLustreFileSystem_DataReadCacheConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#sizing_mode TfLustreFileSystem#sizing_mode}.
	// Experimental.
	SizingMode *string `field:"required" json:"sizingMode" yaml:"sizingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#size TfLustreFileSystem#size}.
	// Experimental.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

