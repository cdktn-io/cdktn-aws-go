package glue


// Experimental.
type AwsSecurityConfiguration_EncryptionConfigurationProperty struct {
	// cloudwatch_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#cloudwatch_encryption AwsSecurityConfiguration#cloudwatch_encryption}
	// Experimental.
	CloudwatchEncryption *AwsSecurityConfiguration_CloudwatchEncryptionProperty `field:"required" json:"cloudwatchEncryption" yaml:"cloudwatchEncryption"`
	// job_bookmarks_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#job_bookmarks_encryption AwsSecurityConfiguration#job_bookmarks_encryption}
	// Experimental.
	JobBookmarksEncryption *AwsSecurityConfiguration_JobBookmarksEncryptionProperty `field:"required" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#s3_encryption AwsSecurityConfiguration#s3_encryption}
	// Experimental.
	S3Encryption *AwsSecurityConfiguration_S3EncryptionProperty `field:"required" json:"s3Encryption" yaml:"s3Encryption"`
}

