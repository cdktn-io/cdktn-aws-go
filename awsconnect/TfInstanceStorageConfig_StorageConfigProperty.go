package awsconnect


// Experimental.
type TfInstanceStorageConfig_StorageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#storage_type TfInstanceStorageConfig#storage_type}.
	// Experimental.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// kinesis_firehose_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_firehose_config TfInstanceStorageConfig#kinesis_firehose_config}
	// Experimental.
	KinesisFirehoseConfig *TfInstanceStorageConfig_KinesisFirehoseConfigProperty `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_stream_config TfInstanceStorageConfig#kinesis_stream_config}
	// Experimental.
	KinesisStreamConfig *TfInstanceStorageConfig_KinesisStreamConfigProperty `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// kinesis_video_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#kinesis_video_stream_config TfInstanceStorageConfig#kinesis_video_stream_config}
	// Experimental.
	KinesisVideoStreamConfig *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#s3_config TfInstanceStorageConfig#s3_config}
	// Experimental.
	S3Config *TfInstanceStorageConfig_S3ConfigProperty `field:"optional" json:"s3Config" yaml:"s3Config"`
}

