package awss3


// Experimental.
type TfBucket_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket TfBucket#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// access_control_translation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#access_control_translation TfBucket#access_control_translation}
	// Experimental.
	AccessControlTranslation *TfBucket_AccessControlTranslationProperty `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#account_id TfBucket#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#metrics TfBucket#metrics}
	// Experimental.
	Metrics *TfBucket_MetricsProperty `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#replica_kms_key_id TfBucket#replica_kms_key_id}.
	// Experimental.
	ReplicaKmsKeyId *string `field:"optional" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
	// replication_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#replication_time TfBucket#replication_time}
	// Experimental.
	ReplicationTime *TfBucket_ReplicationTimeProperty `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#storage_class TfBucket#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

