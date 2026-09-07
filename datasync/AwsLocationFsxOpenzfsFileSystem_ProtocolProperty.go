package datasync


// Experimental.
type AwsLocationFsxOpenzfsFileSystem_ProtocolProperty struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_openzfs_file_system#nfs AwsLocationFsxOpenzfsFileSystem#nfs}
	// Experimental.
	Nfs *AwsLocationFsxOpenzfsFileSystem_NfsProperty `field:"required" json:"nfs" yaml:"nfs"`
}

