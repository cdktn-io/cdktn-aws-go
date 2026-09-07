package eventbridgepipes


// Experimental.
type AwsPipe_CloudwatchLogsParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#log_stream_name AwsPipe#log_stream_name}.
	// Experimental.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#timestamp AwsPipe#timestamp}.
	// Experimental.
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

