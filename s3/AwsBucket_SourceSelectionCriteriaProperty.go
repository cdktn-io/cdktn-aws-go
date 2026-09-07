package s3


// Experimental.
type AwsBucket_SourceSelectionCriteriaProperty struct {
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#sse_kms_encrypted_objects AwsBucket#sse_kms_encrypted_objects}
	// Experimental.
	SseKmsEncryptedObjects *AwsBucket_SseKmsEncryptedObjectsProperty `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

