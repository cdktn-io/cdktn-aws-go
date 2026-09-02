package awseventbridge


// Experimental.
type TfConnection_AuthParametersProperty struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#api_key TfConnection#api_key}
	// Experimental.
	ApiKey *TfConnection_ApiKeyProperty `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#basic TfConnection#basic}
	// Experimental.
	Basic *TfConnection_BasicProperty `field:"optional" json:"basic" yaml:"basic"`
	// connectivity_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#connectivity_parameters TfConnection#connectivity_parameters}
	// Experimental.
	ConnectivityParameters *TfConnection_ConnectivityParametersProperty `field:"optional" json:"connectivityParameters" yaml:"connectivityParameters"`
	// invocation_http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#invocation_http_parameters TfConnection#invocation_http_parameters}
	// Experimental.
	InvocationHttpParameters *TfConnection_InvocationHttpParametersProperty `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#oauth TfConnection#oauth}
	// Experimental.
	Oauth *TfConnection_OauthProperty `field:"optional" json:"oauth" yaml:"oauth"`
}

