package awss3


// Experimental.
type TfBucket_ApplyServerSideEncryptionByDefaultProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#sse_algorithm TfBucket#sse_algorithm}.
	// Experimental.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#kms_master_key_id TfBucket#kms_master_key_id}.
	// Experimental.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

