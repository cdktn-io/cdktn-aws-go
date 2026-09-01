package awseventbridge


// Experimental.
type AwsCloudwatchEventConnection_AuthParametersProperty struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#api_key AwsCloudwatchEventConnection#api_key}
	// Experimental.
	ApiKey *AwsCloudwatchEventConnection_ApiKeyProperty `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#basic AwsCloudwatchEventConnection#basic}
	// Experimental.
	Basic *AwsCloudwatchEventConnection_BasicProperty `field:"optional" json:"basic" yaml:"basic"`
	// connectivity_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#connectivity_parameters AwsCloudwatchEventConnection#connectivity_parameters}
	// Experimental.
	ConnectivityParameters *AwsCloudwatchEventConnection_ConnectivityParametersProperty `field:"optional" json:"connectivityParameters" yaml:"connectivityParameters"`
	// invocation_http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#invocation_http_parameters AwsCloudwatchEventConnection#invocation_http_parameters}
	// Experimental.
	InvocationHttpParameters *AwsCloudwatchEventConnection_InvocationHttpParametersProperty `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#oauth AwsCloudwatchEventConnection#oauth}
	// Experimental.
	Oauth *AwsCloudwatchEventConnection_OauthProperty `field:"optional" json:"oauth" yaml:"oauth"`
}

