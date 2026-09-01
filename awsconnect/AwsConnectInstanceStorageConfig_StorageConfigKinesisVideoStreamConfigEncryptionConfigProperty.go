package awsconnect


// Experimental.
type AwsConnectInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#encryption_type AwsConnectInstanceStorageConfig#encryption_type}.
	// Experimental.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#key_id AwsConnectInstanceStorageConfig#key_id}.
	// Experimental.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

