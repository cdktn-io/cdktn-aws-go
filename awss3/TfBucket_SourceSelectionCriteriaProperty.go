package awss3


// Experimental.
type TfBucket_SourceSelectionCriteriaProperty struct {
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#sse_kms_encrypted_objects TfBucket#sse_kms_encrypted_objects}
	// Experimental.
	SseKmsEncryptedObjects *TfBucket_SseKmsEncryptedObjectsProperty `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

