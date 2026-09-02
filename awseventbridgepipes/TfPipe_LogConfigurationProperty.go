package awseventbridgepipes


// Experimental.
type TfPipe_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#level TfPipe#level}.
	// Experimental.
	Level *string `field:"required" json:"level" yaml:"level"`
	// cloudwatch_logs_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cloudwatch_logs_log_destination TfPipe#cloudwatch_logs_log_destination}
	// Experimental.
	CloudwatchLogsLogDestination *TfPipe_CloudwatchLogsLogDestinationProperty `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// firehose_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#firehose_log_destination TfPipe#firehose_log_destination}
	// Experimental.
	FirehoseLogDestination *TfPipe_FirehoseLogDestinationProperty `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#include_execution_data TfPipe#include_execution_data}.
	// Experimental.
	IncludeExecutionData *[]*string `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// s3_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#s3_log_destination TfPipe#s3_log_destination}
	// Experimental.
	S3LogDestination *TfPipe_S3LogDestinationProperty `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}

