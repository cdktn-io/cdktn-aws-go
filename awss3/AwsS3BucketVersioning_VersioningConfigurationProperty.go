package awss3


// Experimental.
type AwsS3BucketVersioning_VersioningConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_versioning#status AwsS3BucketVersioning#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_versioning#mfa_delete AwsS3BucketVersioning#mfa_delete}.
	// Experimental.
	MfaDelete *string `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

