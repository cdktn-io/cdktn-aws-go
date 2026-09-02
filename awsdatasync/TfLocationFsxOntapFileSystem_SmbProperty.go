package awsdatasync


// Experimental.
type TfLocationFsxOntapFileSystem_SmbProperty struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#mount_options TfLocationFsxOntapFileSystem#mount_options}
	// Experimental.
	MountOptions *TfLocationFsxOntapFileSystem_ProtocolSmbMountOptionsProperty `field:"required" json:"mountOptions" yaml:"mountOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#password TfLocationFsxOntapFileSystem#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#user TfLocationFsxOntapFileSystem#user}.
	// Experimental.
	User *string `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#domain TfLocationFsxOntapFileSystem#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

