package eventbridge


// Experimental.
type AwsConnection_AuthParametersProperty struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#api_key AwsConnection#api_key}
	// Experimental.
	ApiKey *AwsConnection_ApiKeyProperty `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#basic AwsConnection#basic}
	// Experimental.
	Basic *AwsConnection_BasicProperty `field:"optional" json:"basic" yaml:"basic"`
	// connectivity_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#connectivity_parameters AwsConnection#connectivity_parameters}
	// Experimental.
	ConnectivityParameters *AwsConnection_ConnectivityParametersProperty `field:"optional" json:"connectivityParameters" yaml:"connectivityParameters"`
	// invocation_http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#invocation_http_parameters AwsConnection#invocation_http_parameters}
	// Experimental.
	InvocationHttpParameters *AwsConnection_InvocationHttpParametersProperty `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#oauth AwsConnection#oauth}
	// Experimental.
	Oauth *AwsConnection_OauthProperty `field:"optional" json:"oauth" yaml:"oauth"`
}

