package awsconnect


// Experimental.
type AwsConnectInstanceStorageConfig_StorageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#storage_type AwsConnectInstanceStorageConfig#storage_type}.
	// Experimental.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// kinesis_firehose_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_firehose_config AwsConnectInstanceStorageConfig#kinesis_firehose_config}
	// Experimental.
	KinesisFirehoseConfig *AwsConnectInstanceStorageConfig_KinesisFirehoseConfigProperty `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_stream_config AwsConnectInstanceStorageConfig#kinesis_stream_config}
	// Experimental.
	KinesisStreamConfig *AwsConnectInstanceStorageConfig_KinesisStreamConfigProperty `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// kinesis_video_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_video_stream_config AwsConnectInstanceStorageConfig#kinesis_video_stream_config}
	// Experimental.
	KinesisVideoStreamConfig *AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#s3_config AwsConnectInstanceStorageConfig#s3_config}
	// Experimental.
	S3Config *AwsConnectInstanceStorageConfig_S3ConfigProperty `field:"optional" json:"s3Config" yaml:"s3Config"`
}

