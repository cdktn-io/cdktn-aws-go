package s3


// Experimental.
type AwsBucket_ApplyServerSideEncryptionByDefaultProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#sse_algorithm AwsBucket#sse_algorithm}.
	// Experimental.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#kms_master_key_id AwsBucket#kms_master_key_id}.
	// Experimental.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

