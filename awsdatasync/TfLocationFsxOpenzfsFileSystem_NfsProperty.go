package awsdatasync


// Experimental.
type TfLocationFsxOpenzfsFileSystem_NfsProperty struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_openzfs_file_system#mount_options TfLocationFsxOpenzfsFileSystem#mount_options}
	// Experimental.
	MountOptions *TfLocationFsxOpenzfsFileSystem_MountOptionsProperty `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

