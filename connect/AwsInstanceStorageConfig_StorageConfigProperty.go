package connect


// Experimental.
type AwsInstanceStorageConfig_StorageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#storage_type AwsInstanceStorageConfig#storage_type}.
	// Experimental.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// kinesis_firehose_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_firehose_config AwsInstanceStorageConfig#kinesis_firehose_config}
	// Experimental.
	KinesisFirehoseConfig *AwsInstanceStorageConfig_KinesisFirehoseConfigProperty `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_stream_config AwsInstanceStorageConfig#kinesis_stream_config}
	// Experimental.
	KinesisStreamConfig *AwsInstanceStorageConfig_KinesisStreamConfigProperty `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// kinesis_video_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_video_stream_config AwsInstanceStorageConfig#kinesis_video_stream_config}
	// Experimental.
	KinesisVideoStreamConfig *AwsInstanceStorageConfig_KinesisVideoStreamConfigProperty `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#s3_config AwsInstanceStorageConfig#s3_config}
	// Experimental.
	S3Config *AwsInstanceStorageConfig_S3ConfigProperty `field:"optional" json:"s3Config" yaml:"s3Config"`
}

