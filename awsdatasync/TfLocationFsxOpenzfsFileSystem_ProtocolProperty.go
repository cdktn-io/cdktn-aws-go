package awsdatasync


// Experimental.
type TfLocationFsxOpenzfsFileSystem_ProtocolProperty struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_openzfs_file_system#nfs TfLocationFsxOpenzfsFileSystem#nfs}
	// Experimental.
	Nfs *TfLocationFsxOpenzfsFileSystem_NfsProperty `field:"required" json:"nfs" yaml:"nfs"`
}

