package awsfsx


// Experimental.
type TfOpenzfsFileSystem_DiskIopsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#iops TfOpenzfsFileSystem#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#mode TfOpenzfsFileSystem#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

