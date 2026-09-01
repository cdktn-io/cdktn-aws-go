package awss3files


// Experimental.
type AwsS3FilesAccessPoint_CreationPermissionsProperty struct {
	// Owner group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#owner_gid AwsS3FilesAccessPoint#owner_gid}
	// Experimental.
	OwnerGid *float64 `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Owner user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#owner_uid AwsS3FilesAccessPoint#owner_uid}
	// Experimental.
	OwnerUid *float64 `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// POSIX permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#permissions AwsS3FilesAccessPoint#permissions}
	// Experimental.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

