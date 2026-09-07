package kinesisfirehose


// Experimental.
type AwsDeliveryStream_SplunkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_endpoint AwsDeliveryStream#hec_endpoint}.
	// Experimental.
	HecEndpoint *string `field:"required" json:"hecEndpoint" yaml:"hecEndpoint"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration AwsDeliveryStream#s3_configuration}
	// Experimental.
	S3Configuration *AwsDeliveryStream_SplunkConfigurationS3ConfigurationProperty `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#buffering_interval AwsDeliveryStream#buffering_interval}.
	// Experimental.
	BufferingInterval *float64 `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#buffering_size AwsDeliveryStream#buffering_size}.
	// Experimental.
	BufferingSize *float64 `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options AwsDeliveryStream#cloudwatch_logging_options}
	// Experimental.
	CloudwatchLoggingOptions *AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_acknowledgment_timeout AwsDeliveryStream#hec_acknowledgment_timeout}.
	// Experimental.
	HecAcknowledgmentTimeout *float64 `field:"optional" json:"hecAcknowledgmentTimeout" yaml:"hecAcknowledgmentTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_endpoint_type AwsDeliveryStream#hec_endpoint_type}.
	// Experimental.
	HecEndpointType *string `field:"optional" json:"hecEndpointType" yaml:"hecEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hec_token AwsDeliveryStream#hec_token}.
	// Experimental.
	HecToken *string `field:"optional" json:"hecToken" yaml:"hecToken"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration AwsDeliveryStream#processing_configuration}
	// Experimental.
	ProcessingConfiguration *AwsDeliveryStream_SplunkConfigurationProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration AwsDeliveryStream#retry_duration}.
	// Experimental.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode AwsDeliveryStream#s3_backup_mode}.
	// Experimental.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration AwsDeliveryStream#secrets_manager_configuration}
	// Experimental.
	SecretsManagerConfiguration *AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
}

