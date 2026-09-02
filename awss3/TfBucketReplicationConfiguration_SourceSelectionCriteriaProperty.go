package awss3


// Experimental.
type TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty struct {
	// replica_modifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replica_modifications TfBucketReplicationConfiguration#replica_modifications}
	// Experimental.
	ReplicaModifications *TfBucketReplicationConfiguration_ReplicaModificationsProperty `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#sse_kms_encrypted_objects TfBucketReplicationConfiguration#sse_kms_encrypted_objects}
	// Experimental.
	SseKmsEncryptedObjects *TfBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

