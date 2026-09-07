package datasync


// Experimental.
type AwsLocationFsxOntapFileSystem_ProtocolProperty struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#nfs AwsLocationFsxOntapFileSystem#nfs}
	// Experimental.
	Nfs *AwsLocationFsxOntapFileSystem_NfsProperty `field:"optional" json:"nfs" yaml:"nfs"`
	// smb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#smb AwsLocationFsxOntapFileSystem#smb}
	// Experimental.
	Smb *AwsLocationFsxOntapFileSystem_SmbProperty `field:"optional" json:"smb" yaml:"smb"`
}

