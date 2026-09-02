package awsfsx


// Experimental.
type TfLustreFileSystem_MetadataConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#iops TfLustreFileSystem#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#mode TfLustreFileSystem#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

