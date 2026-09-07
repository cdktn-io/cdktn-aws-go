package vpnsitetosite


// Experimental.
type AwsConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#bgp_log_enabled AwsConnection#bgp_log_enabled}.
	// Experimental.
	BgpLogEnabled interface{} `field:"optional" json:"bgpLogEnabled" yaml:"bgpLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#bgp_log_group_arn AwsConnection#bgp_log_group_arn}.
	// Experimental.
	BgpLogGroupArn *string `field:"optional" json:"bgpLogGroupArn" yaml:"bgpLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#bgp_log_output_format AwsConnection#bgp_log_output_format}.
	// Experimental.
	BgpLogOutputFormat *string `field:"optional" json:"bgpLogOutputFormat" yaml:"bgpLogOutputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#log_enabled AwsConnection#log_enabled}.
	// Experimental.
	LogEnabled interface{} `field:"optional" json:"logEnabled" yaml:"logEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#log_group_arn AwsConnection#log_group_arn}.
	// Experimental.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#log_output_format AwsConnection#log_output_format}.
	// Experimental.
	LogOutputFormat *string `field:"optional" json:"logOutputFormat" yaml:"logOutputFormat"`
}

