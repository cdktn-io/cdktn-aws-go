package awss3tables


// Experimental.
type TfTableBucket_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#kms_key_arn TfTableBucket#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#sse_algorithm TfTableBucket#sse_algorithm}.
	// Experimental.
	SseAlgorithm *string `field:"optional" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}

