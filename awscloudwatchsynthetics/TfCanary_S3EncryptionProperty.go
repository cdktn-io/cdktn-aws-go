package awscloudwatchsynthetics


// Experimental.
type TfCanary_S3EncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#encryption_mode TfCanary#encryption_mode}.
	// Experimental.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#kms_key_arn TfCanary#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

