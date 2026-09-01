package awss3


// Experimental.
type AwsS3BucketInventory_EncryptionProperty struct {
	// sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#sse_kms AwsS3BucketInventory#sse_kms}
	// Experimental.
	SseKms *AwsS3BucketInventory_SseKmsProperty `field:"optional" json:"sseKms" yaml:"sseKms"`
	// sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#sse_s3 AwsS3BucketInventory#sse_s3}
	// Experimental.
	SseS3 *AwsS3BucketInventory_SseS3Property `field:"optional" json:"sseS3" yaml:"sseS3"`
}

