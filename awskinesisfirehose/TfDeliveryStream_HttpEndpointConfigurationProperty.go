package awskinesisfirehose


// Experimental.
type TfDeliveryStream_HttpEndpointConfigurationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration TfDeliveryStream#s3_configuration}
	// Experimental.
	S3Configuration *TfDeliveryStream_HttpEndpointConfigurationS3ConfigurationProperty `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#url TfDeliveryStream#url}.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#access_key TfDeliveryStream#access_key}.
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
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
	CloudwatchLoggingOptions *TfDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsProperty `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#name TfDeliveryStream#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration TfDeliveryStream#processing_configuration}
	// Experimental.
	ProcessingConfiguration *TfDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// request_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#request_configuration TfDeliveryStream#request_configuration}
	// Experimental.
	RequestConfiguration *TfDeliveryStream_RequestConfigurationProperty `field:"optional" json:"requestConfiguration" yaml:"requestConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration TfDeliveryStream#retry_duration}.
	// Experimental.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#role_arn TfDeliveryStream#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode TfDeliveryStream#s3_backup_mode}.
	// Experimental.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration TfDeliveryStream#secrets_manager_configuration}
	// Experimental.
	SecretsManagerConfiguration *TfDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationProperty `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
}

