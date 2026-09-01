package awsconnect


// Experimental.
type AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty struct {
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#encryption_config AwsConnectInstanceStorageConfig#encryption_config}
	// Experimental.
	EncryptionConfig *AwsConnectInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty `field:"required" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#prefix AwsConnectInstanceStorageConfig#prefix}.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#retention_period_hours AwsConnectInstanceStorageConfig#retention_period_hours}.
	// Experimental.
	RetentionPeriodHours *float64 `field:"required" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
}

