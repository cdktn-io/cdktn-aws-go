package vpnsitetosite


// Experimental.
type AwsConnection_Tunnel2LogOptionsProperty struct {
	// cloudwatch_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#cloudwatch_log_options AwsConnection#cloudwatch_log_options}
	// Experimental.
	CloudwatchLogOptions *AwsConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

