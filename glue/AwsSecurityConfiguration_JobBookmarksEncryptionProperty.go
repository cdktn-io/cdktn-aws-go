package glue


// Experimental.
type AwsSecurityConfiguration_JobBookmarksEncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#job_bookmarks_encryption_mode AwsSecurityConfiguration#job_bookmarks_encryption_mode}.
	// Experimental.
	JobBookmarksEncryptionMode *string `field:"optional" json:"jobBookmarksEncryptionMode" yaml:"jobBookmarksEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#kms_key_arn AwsSecurityConfiguration#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

