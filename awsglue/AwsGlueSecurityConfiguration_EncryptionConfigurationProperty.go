package awsglue


// Experimental.
type AwsGlueSecurityConfiguration_EncryptionConfigurationProperty struct {
	// cloudwatch_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#cloudwatch_encryption AwsGlueSecurityConfiguration#cloudwatch_encryption}
	// Experimental.
	CloudwatchEncryption *AwsGlueSecurityConfiguration_CloudwatchEncryptionProperty `field:"required" json:"cloudwatchEncryption" yaml:"cloudwatchEncryption"`
	// job_bookmarks_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#job_bookmarks_encryption AwsGlueSecurityConfiguration#job_bookmarks_encryption}
	// Experimental.
	JobBookmarksEncryption *AwsGlueSecurityConfiguration_JobBookmarksEncryptionProperty `field:"required" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#s3_encryption AwsGlueSecurityConfiguration#s3_encryption}
	// Experimental.
	S3Encryption *AwsGlueSecurityConfiguration_S3EncryptionProperty `field:"required" json:"s3Encryption" yaml:"s3Encryption"`
}

