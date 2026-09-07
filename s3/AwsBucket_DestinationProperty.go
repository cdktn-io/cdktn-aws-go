package s3


// Experimental.
type AwsBucket_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket AwsBucket#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// access_control_translation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#access_control_translation AwsBucket#access_control_translation}
	// Experimental.
	AccessControlTranslation *AwsBucket_AccessControlTranslationProperty `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#account_id AwsBucket#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#metrics AwsBucket#metrics}
	// Experimental.
	Metrics *AwsBucket_MetricsProperty `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#replica_kms_key_id AwsBucket#replica_kms_key_id}.
	// Experimental.
	ReplicaKmsKeyId *string `field:"optional" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
	// replication_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#replication_time AwsBucket#replication_time}
	// Experimental.
	ReplicationTime *AwsBucket_ReplicationTimeProperty `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#storage_class AwsBucket#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

