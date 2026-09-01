package awsglue


// Experimental.
type AwsGlueSecurityConfiguration_CloudwatchEncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#cloudwatch_encryption_mode AwsGlueSecurityConfiguration#cloudwatch_encryption_mode}.
	// Experimental.
	CloudwatchEncryptionMode *string `field:"optional" json:"cloudwatchEncryptionMode" yaml:"cloudwatchEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_security_configuration#kms_key_arn AwsGlueSecurityConfiguration#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

