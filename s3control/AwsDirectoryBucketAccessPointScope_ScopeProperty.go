package s3control


// Experimental.
type AwsDirectoryBucketAccessPointScope_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_directory_bucket_access_point_scope#permissions AwsDirectoryBucketAccessPointScope#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_directory_bucket_access_point_scope#prefixes AwsDirectoryBucketAccessPointScope#prefixes}.
	// Experimental.
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}

