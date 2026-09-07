package datasync


// Experimental.
type AwsLocationFsxOntapFileSystem_SmbProperty struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#mount_options AwsLocationFsxOntapFileSystem#mount_options}
	// Experimental.
	MountOptions *AwsLocationFsxOntapFileSystem_ProtocolSmbMountOptionsProperty `field:"required" json:"mountOptions" yaml:"mountOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#password AwsLocationFsxOntapFileSystem#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#user AwsLocationFsxOntapFileSystem#user}.
	// Experimental.
	User *string `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_fsx_ontap_file_system#domain AwsLocationFsxOntapFileSystem#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

