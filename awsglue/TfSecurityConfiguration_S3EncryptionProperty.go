package awsglue


// Experimental.
type TfSecurityConfiguration_S3EncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#kms_key_arn TfSecurityConfiguration#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#s3_encryption_mode TfSecurityConfiguration#s3_encryption_mode}.
	// Experimental.
	S3EncryptionMode *string `field:"optional" json:"s3EncryptionMode" yaml:"s3EncryptionMode"`
}

