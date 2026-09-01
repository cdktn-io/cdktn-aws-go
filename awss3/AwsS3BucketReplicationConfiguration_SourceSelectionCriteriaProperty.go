package awss3


// Experimental.
type AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty struct {
	// replica_modifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replica_modifications AwsS3BucketReplicationConfiguration#replica_modifications}
	// Experimental.
	ReplicaModifications *AwsS3BucketReplicationConfiguration_ReplicaModificationsProperty `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#sse_kms_encrypted_objects AwsS3BucketReplicationConfiguration#sse_kms_encrypted_objects}
	// Experimental.
	SseKmsEncryptedObjects *AwsS3BucketReplicationConfiguration_SseKmsEncryptedObjectsProperty `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

