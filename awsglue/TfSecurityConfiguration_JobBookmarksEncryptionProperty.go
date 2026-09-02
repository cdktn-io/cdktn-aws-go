package awsglue


// Experimental.
type TfSecurityConfiguration_JobBookmarksEncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#job_bookmarks_encryption_mode TfSecurityConfiguration#job_bookmarks_encryption_mode}.
	// Experimental.
	JobBookmarksEncryptionMode *string `field:"optional" json:"jobBookmarksEncryptionMode" yaml:"jobBookmarksEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#kms_key_arn TfSecurityConfiguration#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

