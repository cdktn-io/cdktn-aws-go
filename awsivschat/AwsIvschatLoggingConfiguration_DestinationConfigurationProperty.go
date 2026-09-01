package awsivschat


// Experimental.
type AwsIvschatLoggingConfiguration_DestinationConfigurationProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#cloudwatch_logs AwsIvschatLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsIvschatLoggingConfiguration_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#firehose AwsIvschatLoggingConfiguration#firehose}
	// Experimental.
	Firehose *AwsIvschatLoggingConfiguration_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#s3 AwsIvschatLoggingConfiguration#s3}
	// Experimental.
	S3 *AwsIvschatLoggingConfiguration_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

