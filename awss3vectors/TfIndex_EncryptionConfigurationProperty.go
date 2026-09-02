package awss3vectors


// Experimental.
type TfIndex_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3vectors_index#kms_key_arn TfIndex#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3vectors_index#sse_type TfIndex#sse_type}.
	// Experimental.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

