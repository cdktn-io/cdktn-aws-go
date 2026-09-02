package awsdatasync


// Experimental.
type TfLocationFsxOntapFileSystem_ProtocolProperty struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#nfs TfLocationFsxOntapFileSystem#nfs}
	// Experimental.
	Nfs *TfLocationFsxOntapFileSystem_NfsProperty `field:"optional" json:"nfs" yaml:"nfs"`
	// smb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#smb TfLocationFsxOntapFileSystem#smb}
	// Experimental.
	Smb *TfLocationFsxOntapFileSystem_SmbProperty `field:"optional" json:"smb" yaml:"smb"`
}

