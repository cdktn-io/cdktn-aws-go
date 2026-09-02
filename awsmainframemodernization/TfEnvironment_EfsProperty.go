package awsmainframemodernization


// Experimental.
type TfEnvironment_EfsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/m2_environment#file_system_id TfEnvironment#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/m2_environment#mount_point TfEnvironment#mount_point}.
	// Experimental.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}

