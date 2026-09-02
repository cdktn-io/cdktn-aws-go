package awsfsx


// Experimental.
type TfWindowsFileSystem_DiskIopsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#iops TfWindowsFileSystem#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#mode TfWindowsFileSystem#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

