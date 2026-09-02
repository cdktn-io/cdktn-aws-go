package awsverifiedaccess


// Experimental.
type TfInstanceLoggingConfiguration_AccessLogsProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#cloudwatch_logs TfInstanceLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfInstanceLoggingConfiguration_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#include_trust_context TfInstanceLoggingConfiguration#include_trust_context}.
	// Experimental.
	IncludeTrustContext interface{} `field:"optional" json:"includeTrustContext" yaml:"includeTrustContext"`
	// kinesis_data_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#kinesis_data_firehose TfInstanceLoggingConfiguration#kinesis_data_firehose}
	// Experimental.
	KinesisDataFirehose *TfInstanceLoggingConfiguration_KinesisDataFirehoseProperty `field:"optional" json:"kinesisDataFirehose" yaml:"kinesisDataFirehose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#log_version TfInstanceLoggingConfiguration#log_version}.
	// Experimental.
	LogVersion *string `field:"optional" json:"logVersion" yaml:"logVersion"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#s3 TfInstanceLoggingConfiguration#s3}
	// Experimental.
	S3 *TfInstanceLoggingConfiguration_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

