package eventbridgepipes


// Experimental.
type AwsPipe_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#level AwsPipe#level}.
	// Experimental.
	Level *string `field:"required" json:"level" yaml:"level"`
	// cloudwatch_logs_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cloudwatch_logs_log_destination AwsPipe#cloudwatch_logs_log_destination}
	// Experimental.
	CloudwatchLogsLogDestination *AwsPipe_CloudwatchLogsLogDestinationProperty `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// firehose_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#firehose_log_destination AwsPipe#firehose_log_destination}
	// Experimental.
	FirehoseLogDestination *AwsPipe_FirehoseLogDestinationProperty `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#include_execution_data AwsPipe#include_execution_data}.
	// Experimental.
	IncludeExecutionData *[]*string `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// s3_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#s3_log_destination AwsPipe#s3_log_destination}
	// Experimental.
	S3LogDestination *AwsPipe_S3LogDestinationProperty `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}

