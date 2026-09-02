package awsvpnclient


// Experimental.
type TfEndpoint_ConnectionLogOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#enabled TfEndpoint#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#cloudwatch_log_group TfEndpoint#cloudwatch_log_group}.
	// Experimental.
	CloudwatchLogGroup *string `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#cloudwatch_log_stream TfEndpoint#cloudwatch_log_stream}.
	// Experimental.
	CloudwatchLogStream *string `field:"optional" json:"cloudwatchLogStream" yaml:"cloudwatchLogStream"`
}

