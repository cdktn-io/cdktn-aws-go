package awsfsx


// Experimental.
type AwsFsxS3AccessPointAttachment_FileSystemIdentityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#type AwsFsxS3AccessPointAttachment#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// posix_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#posix_user AwsFsxS3AccessPointAttachment#posix_user}
	// Experimental.
	PosixUser interface{} `field:"optional" json:"posixUser" yaml:"posixUser"`
}

