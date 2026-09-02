package awss3


// Experimental.
type TfBucketMetadataConfiguration_MetadataConfigurationInventoryTableConfigurationEncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#sse_algorithm TfBucketMetadataConfiguration#sse_algorithm}.
	// Experimental.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#kms_key_arn TfBucketMetadataConfiguration#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

