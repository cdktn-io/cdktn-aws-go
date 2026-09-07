package s3


// Experimental.
type AwsBucketInventory_EncryptionProperty struct {
	// sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#sse_kms AwsBucketInventory#sse_kms}
	// Experimental.
	SseKms *AwsBucketInventory_SseKmsProperty `field:"optional" json:"sseKms" yaml:"sseKms"`
	// sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#sse_s3 AwsBucketInventory#sse_s3}
	// Experimental.
	SseS3 *AwsBucketInventory_SseS3Property `field:"optional" json:"sseS3" yaml:"sseS3"`
}

