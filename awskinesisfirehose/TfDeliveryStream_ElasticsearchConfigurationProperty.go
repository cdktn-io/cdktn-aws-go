package awskinesisfirehose


// Experimental.
type TfDeliveryStream_ElasticsearchConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#index_name TfDeliveryStream#index_name}.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#role_arn TfDeliveryStream#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration TfDeliveryStream#s3_configuration}
	// Experimental.
	S3Configuration *TfDeliveryStream_ElasticsearchConfigurationS3ConfigurationProperty `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
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
	CloudwatchLoggingOptions *TfDeliveryStream_ElasticsearchConfigurationCloudwatchLoggingOptionsProperty `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cluster_endpoint TfDeliveryStream#cluster_endpoint}.
	// Experimental.
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#domain_arn TfDeliveryStream#domain_arn}.
	// Experimental.
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#index_rotation_period TfDeliveryStream#index_rotation_period}.
	// Experimental.
	IndexRotationPeriod *string `field:"optional" json:"indexRotationPeriod" yaml:"indexRotationPeriod"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration TfDeliveryStream#processing_configuration}
	// Experimental.
	ProcessingConfiguration *TfDeliveryStream_ElasticsearchConfigurationProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration TfDeliveryStream#retry_duration}.
	// Experimental.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode TfDeliveryStream#s3_backup_mode}.
	// Experimental.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#type_name TfDeliveryStream#type_name}.
	// Experimental.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#vpc_config TfDeliveryStream#vpc_config}
	// Experimental.
	VpcConfig *TfDeliveryStream_ElasticsearchConfigurationVpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

