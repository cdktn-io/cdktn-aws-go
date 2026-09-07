package s3vectors


// Experimental.
type AwsVectorBucket_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3vectors_vector_bucket#kms_key_arn AwsVectorBucket#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3vectors_vector_bucket#sse_type AwsVectorBucket#sse_type}.
	// Experimental.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

