package awskinesisfirehose


// Experimental.
type TfDeliveryStream_SplunkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_endpoint TfDeliveryStream#hec_endpoint}.
	// Experimental.
	HecEndpoint *string `field:"required" json:"hecEndpoint" yaml:"hecEndpoint"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration TfDeliveryStream#s3_configuration}
	// Experimental.
	S3Configuration *TfDeliveryStream_SplunkConfigurationS3ConfigurationProperty `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#buffering_interval TfDeliveryStream#buffering_interval}.
	// Experimental.
	BufferingInterval *float64 `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#buffering_size TfDeliveryStream#buffering_size}.
	// Experimental.
	BufferingSize *float64 `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options TfDeliveryStream#cloudwatch_logging_options}
	// Experimental.
	CloudwatchLoggingOptions *TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_acknowledgment_timeout TfDeliveryStream#hec_acknowledgment_timeout}.
	// Experimental.
	HecAcknowledgmentTimeout *float64 `field:"optional" json:"hecAcknowledgmentTimeout" yaml:"hecAcknowledgmentTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_endpoint_type TfDeliveryStream#hec_endpoint_type}.
	// Experimental.
	HecEndpointType *string `field:"optional" json:"hecEndpointType" yaml:"hecEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_token TfDeliveryStream#hec_token}.
	// Experimental.
	HecToken *string `field:"optional" json:"hecToken" yaml:"hecToken"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration TfDeliveryStream#processing_configuration}
	// Experimental.
	ProcessingConfiguration *TfDeliveryStream_SplunkConfigurationProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration TfDeliveryStream#retry_duration}.
	// Experimental.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode TfDeliveryStream#s3_backup_mode}.
	// Experimental.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration TfDeliveryStream#secrets_manager_configuration}
	// Experimental.
	SecretsManagerConfiguration *TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
}

