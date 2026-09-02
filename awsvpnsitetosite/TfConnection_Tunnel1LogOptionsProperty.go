package awsvpnsitetosite


// Experimental.
type TfConnection_Tunnel1LogOptionsProperty struct {
	// cloudwatch_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#cloudwatch_log_options TfConnection#cloudwatch_log_options}
	// Experimental.
	CloudwatchLogOptions *TfConnection_Tunnel1LogOptionsCloudwatchLogOptionsProperty `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

