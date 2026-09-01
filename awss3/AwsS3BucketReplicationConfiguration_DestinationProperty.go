package awss3


// Experimental.
type AwsS3BucketReplicationConfiguration_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#bucket AwsS3BucketReplicationConfiguration#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// access_control_translation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#access_control_translation AwsS3BucketReplicationConfiguration#access_control_translation}
	// Experimental.
	AccessControlTranslation *AwsS3BucketReplicationConfiguration_AccessControlTranslationProperty `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#account AwsS3BucketReplicationConfiguration#account}.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#encryption_configuration AwsS3BucketReplicationConfiguration#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *AwsS3BucketReplicationConfiguration_EncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#metrics AwsS3BucketReplicationConfiguration#metrics}
	// Experimental.
	Metrics *AwsS3BucketReplicationConfiguration_MetricsProperty `field:"optional" json:"metrics" yaml:"metrics"`
	// replication_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#replication_time AwsS3BucketReplicationConfiguration#replication_time}
	// Experimental.
	ReplicationTime *AwsS3BucketReplicationConfiguration_ReplicationTimeProperty `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#storage_class AwsS3BucketReplicationConfiguration#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

