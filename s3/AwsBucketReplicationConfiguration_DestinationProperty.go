package s3


// Experimental.
type AwsBucketReplicationConfiguration_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#bucket AwsBucketReplicationConfiguration#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// access_control_translation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#access_control_translation AwsBucketReplicationConfiguration#access_control_translation}
	// Experimental.
	AccessControlTranslation *AwsBucketReplicationConfiguration_AccessControlTranslationProperty `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#account AwsBucketReplicationConfiguration#account}.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#encryption_configuration AwsBucketReplicationConfiguration#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *AwsBucketReplicationConfiguration_EncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#metrics AwsBucketReplicationConfiguration#metrics}
	// Experimental.
	Metrics *AwsBucketReplicationConfiguration_MetricsProperty `field:"optional" json:"metrics" yaml:"metrics"`
	// replication_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replication_time AwsBucketReplicationConfiguration#replication_time}
	// Experimental.
	ReplicationTime *AwsBucketReplicationConfiguration_ReplicationTimeProperty `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#storage_class AwsBucketReplicationConfiguration#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

