package awss3files


// Experimental.
type TfAccessPoint_PosixUserProperty struct {
	// POSIX group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#gid TfAccessPoint#gid}
	// Experimental.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// POSIX user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#uid TfAccessPoint#uid}
	// Experimental.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#secondary_gids TfAccessPoint#secondary_gids}
	// Experimental.
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

