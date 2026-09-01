package awss3control


// Experimental.
type AwsS3ControlDirectoryBucketAccessPointScope_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_directory_bucket_access_point_scope#permissions AwsS3ControlDirectoryBucketAccessPointScope#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_directory_bucket_access_point_scope#prefixes AwsS3ControlDirectoryBucketAccessPointScope#prefixes}.
	// Experimental.
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}

