package awss3


// Experimental.
type AwsS3Bucket_ServerSideEncryptionConfigurationProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rule AwsS3Bucket#rule}
	// Experimental.
	Rule *AwsS3Bucket_ServerSideEncryptionConfigurationRuleProperty `field:"required" json:"rule" yaml:"rule"`
}

