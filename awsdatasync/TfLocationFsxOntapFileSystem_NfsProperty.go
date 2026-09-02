package awsdatasync


// Experimental.
type TfLocationFsxOntapFileSystem_NfsProperty struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#mount_options TfLocationFsxOntapFileSystem#mount_options}
	// Experimental.
	MountOptions *TfLocationFsxOntapFileSystem_ProtocolNfsMountOptionsProperty `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

