package s3


// Experimental.
type AwsBucketReplicationConfiguration_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replica_kms_key_id AwsBucketReplicationConfiguration#replica_kms_key_id}.
	// Experimental.
	ReplicaKmsKeyId *string `field:"required" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
}

