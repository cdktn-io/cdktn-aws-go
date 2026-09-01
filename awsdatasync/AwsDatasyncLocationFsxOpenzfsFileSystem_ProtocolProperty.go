package awsdatasync


// Experimental.
type AwsDatasyncLocationFsxOpenzfsFileSystem_ProtocolProperty struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_openzfs_file_system#nfs AwsDatasyncLocationFsxOpenzfsFileSystem#nfs}
	// Experimental.
	Nfs *AwsDatasyncLocationFsxOpenzfsFileSystem_NfsProperty `field:"required" json:"nfs" yaml:"nfs"`
}

