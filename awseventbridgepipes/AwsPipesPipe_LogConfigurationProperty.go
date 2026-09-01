package awseventbridgepipes


// Experimental.
type AwsPipesPipe_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#level AwsPipesPipe#level}.
	// Experimental.
	Level *string `field:"required" json:"level" yaml:"level"`
	// cloudwatch_logs_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cloudwatch_logs_log_destination AwsPipesPipe#cloudwatch_logs_log_destination}
	// Experimental.
	CloudwatchLogsLogDestination *AwsPipesPipe_CloudwatchLogsLogDestinationProperty `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// firehose_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#firehose_log_destination AwsPipesPipe#firehose_log_destination}
	// Experimental.
	FirehoseLogDestination *AwsPipesPipe_FirehoseLogDestinationProperty `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#include_execution_data AwsPipesPipe#include_execution_data}.
	// Experimental.
	IncludeExecutionData *[]*string `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// s3_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#s3_log_destination AwsPipesPipe#s3_log_destination}
	// Experimental.
	S3LogDestination *AwsPipesPipe_S3LogDestinationProperty `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}

