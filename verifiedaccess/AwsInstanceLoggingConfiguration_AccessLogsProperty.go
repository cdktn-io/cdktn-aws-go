package verifiedaccess


// Experimental.
type AwsInstanceLoggingConfiguration_AccessLogsProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#cloudwatch_logs AwsInstanceLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsInstanceLoggingConfiguration_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#include_trust_context AwsInstanceLoggingConfiguration#include_trust_context}.
	// Experimental.
	IncludeTrustContext interface{} `field:"optional" json:"includeTrustContext" yaml:"includeTrustContext"`
	// kinesis_data_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#kinesis_data_firehose AwsInstanceLoggingConfiguration#kinesis_data_firehose}
	// Experimental.
	KinesisDataFirehose *AwsInstanceLoggingConfiguration_KinesisDataFirehoseProperty `field:"optional" json:"kinesisDataFirehose" yaml:"kinesisDataFirehose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#log_version AwsInstanceLoggingConfiguration#log_version}.
	// Experimental.
	LogVersion *string `field:"optional" json:"logVersion" yaml:"logVersion"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#s3 AwsInstanceLoggingConfiguration#s3}
	// Experimental.
	S3 *AwsInstanceLoggingConfiguration_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

