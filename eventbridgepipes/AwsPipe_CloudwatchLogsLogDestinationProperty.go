package eventbridgepipes


// Experimental.
type AwsPipe_CloudwatchLogsLogDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#log_group_arn AwsPipe#log_group_arn}.
	// Experimental.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

