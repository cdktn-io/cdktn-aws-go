package awsconnect


// Experimental.
type TfInstanceStorageConfig_KinesisVideoStreamConfigProperty struct {
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#encryption_config TfInstanceStorageConfig#encryption_config}
	// Experimental.
	EncryptionConfig *TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty `field:"required" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#prefix TfInstanceStorageConfig#prefix}.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#retention_period_hours TfInstanceStorageConfig#retention_period_hours}.
	// Experimental.
	RetentionPeriodHours *float64 `field:"required" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
}

