package awsvpnsitetosite


// Experimental.
type AwsVpnConnection_Tunnel1LogOptionsCloudwatchLogOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#bgp_log_enabled AwsVpnConnection#bgp_log_enabled}.
	// Experimental.
	BgpLogEnabled interface{} `field:"optional" json:"bgpLogEnabled" yaml:"bgpLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#bgp_log_group_arn AwsVpnConnection#bgp_log_group_arn}.
	// Experimental.
	BgpLogGroupArn *string `field:"optional" json:"bgpLogGroupArn" yaml:"bgpLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#bgp_log_output_format AwsVpnConnection#bgp_log_output_format}.
	// Experimental.
	BgpLogOutputFormat *string `field:"optional" json:"bgpLogOutputFormat" yaml:"bgpLogOutputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#log_enabled AwsVpnConnection#log_enabled}.
	// Experimental.
	LogEnabled interface{} `field:"optional" json:"logEnabled" yaml:"logEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#log_group_arn AwsVpnConnection#log_group_arn}.
	// Experimental.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#log_output_format AwsVpnConnection#log_output_format}.
	// Experimental.
	LogOutputFormat *string `field:"optional" json:"logOutputFormat" yaml:"logOutputFormat"`
}

