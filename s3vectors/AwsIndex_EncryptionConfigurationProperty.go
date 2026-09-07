package s3vectors


// Experimental.
type AwsIndex_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3vectors_index#kms_key_arn AwsIndex#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3vectors_index#sse_type AwsIndex#sse_type}.
	// Experimental.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

