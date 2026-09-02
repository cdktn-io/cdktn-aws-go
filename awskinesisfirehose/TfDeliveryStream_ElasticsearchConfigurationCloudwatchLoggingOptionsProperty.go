package awskinesisfirehose


// Experimental.
type TfDeliveryStream_ElasticsearchConfigurationCloudwatchLoggingOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enabled TfDeliveryStream#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#log_group_name TfDeliveryStream#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#log_stream_name TfDeliveryStream#log_stream_name}.
	// Experimental.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

