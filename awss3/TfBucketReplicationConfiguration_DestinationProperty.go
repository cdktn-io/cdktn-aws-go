package awss3


// Experimental.
type TfBucketReplicationConfiguration_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#bucket TfBucketReplicationConfiguration#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// access_control_translation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#access_control_translation TfBucketReplicationConfiguration#access_control_translation}
	// Experimental.
	AccessControlTranslation *TfBucketReplicationConfiguration_AccessControlTranslationProperty `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#account TfBucketReplicationConfiguration#account}.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#encryption_configuration TfBucketReplicationConfiguration#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *TfBucketReplicationConfiguration_EncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#metrics TfBucketReplicationConfiguration#metrics}
	// Experimental.
	Metrics *TfBucketReplicationConfiguration_MetricsProperty `field:"optional" json:"metrics" yaml:"metrics"`
	// replication_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replication_time TfBucketReplicationConfiguration#replication_time}
	// Experimental.
	ReplicationTime *TfBucketReplicationConfiguration_ReplicationTimeProperty `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#storage_class TfBucketReplicationConfiguration#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

