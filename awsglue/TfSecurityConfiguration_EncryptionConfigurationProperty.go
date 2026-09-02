package awsglue


// Experimental.
type TfSecurityConfiguration_EncryptionConfigurationProperty struct {
	// cloudwatch_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#cloudwatch_encryption TfSecurityConfiguration#cloudwatch_encryption}
	// Experimental.
	CloudwatchEncryption *TfSecurityConfiguration_CloudwatchEncryptionProperty `field:"required" json:"cloudwatchEncryption" yaml:"cloudwatchEncryption"`
	// job_bookmarks_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#job_bookmarks_encryption TfSecurityConfiguration#job_bookmarks_encryption}
	// Experimental.
	JobBookmarksEncryption *TfSecurityConfiguration_JobBookmarksEncryptionProperty `field:"required" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#s3_encryption TfSecurityConfiguration#s3_encryption}
	// Experimental.
	S3Encryption *TfSecurityConfiguration_S3EncryptionProperty `field:"required" json:"s3Encryption" yaml:"s3Encryption"`
}

