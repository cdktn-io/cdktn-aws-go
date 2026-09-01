package awsconnect


// Experimental.
type AwsConnectInstanceStorageConfig_S3ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#bucket_name AwsConnectInstanceStorageConfig#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#bucket_prefix AwsConnectInstanceStorageConfig#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#encryption_config AwsConnectInstanceStorageConfig#encryption_config}
	// Experimental.
	EncryptionConfig *AwsConnectInstanceStorageConfig_StorageConfigS3ConfigEncryptionConfigProperty `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

