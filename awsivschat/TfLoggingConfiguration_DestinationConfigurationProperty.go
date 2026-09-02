package awsivschat


// Experimental.
type TfLoggingConfiguration_DestinationConfigurationProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#cloudwatch_logs TfLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfLoggingConfiguration_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#firehose TfLoggingConfiguration#firehose}
	// Experimental.
	Firehose *TfLoggingConfiguration_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#s3 TfLoggingConfiguration#s3}
	// Experimental.
	S3 *TfLoggingConfiguration_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

