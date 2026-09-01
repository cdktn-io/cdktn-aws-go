package awsdatasync


// Experimental.
type AwsDatasyncLocationFsxOntapFileSystem_ProtocolProperty struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#nfs AwsDatasyncLocationFsxOntapFileSystem#nfs}
	// Experimental.
	Nfs *AwsDatasyncLocationFsxOntapFileSystem_NfsProperty `field:"optional" json:"nfs" yaml:"nfs"`
	// smb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#smb AwsDatasyncLocationFsxOntapFileSystem#smb}
	// Experimental.
	Smb *AwsDatasyncLocationFsxOntapFileSystem_SmbProperty `field:"optional" json:"smb" yaml:"smb"`
}

