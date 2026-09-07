package ivschat


// Experimental.
type AwsLoggingConfiguration_DestinationConfigurationProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#cloudwatch_logs AwsLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsLoggingConfiguration_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#firehose AwsLoggingConfiguration#firehose}
	// Experimental.
	Firehose *AwsLoggingConfiguration_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivschat_logging_configuration#s3 AwsLoggingConfiguration#s3}
	// Experimental.
	S3 *AwsLoggingConfiguration_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

