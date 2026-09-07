package datasync


// Experimental.
type AwsLocationFsxOntapFileSystem_NfsProperty struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#mount_options AwsLocationFsxOntapFileSystem#mount_options}
	// Experimental.
	MountOptions *AwsLocationFsxOntapFileSystem_ProtocolNfsMountOptionsProperty `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

