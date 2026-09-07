package s3


// Experimental.
type AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty struct {
	// replica_modifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replica_modifications AwsBucketReplicationConfiguration#replica_modifications}
	// Experimental.
	ReplicaModifications *AwsBucketReplicationConfiguration_ReplicaModificationsProperty `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#sse_kms_encrypted_objects AwsBucketReplicationConfiguration#sse_kms_encrypted_objects}
	// Experimental.
	SseKmsEncryptedObjects *AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

