package awseventbridge


// Experimental.
type AwsCloudwatchEventConnection_OauthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#authorization_endpoint AwsCloudwatchEventConnection#authorization_endpoint}.
	// Experimental.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#http_method AwsCloudwatchEventConnection#http_method}.
	// Experimental.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// oauth_http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#oauth_http_parameters AwsCloudwatchEventConnection#oauth_http_parameters}
	// Experimental.
	OauthHttpParameters *AwsCloudwatchEventConnection_OauthHttpParametersProperty `field:"required" json:"oauthHttpParameters" yaml:"oauthHttpParameters"`
	// client_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#client_parameters AwsCloudwatchEventConnection#client_parameters}
	// Experimental.
	ClientParameters *AwsCloudwatchEventConnection_ClientParametersProperty `field:"optional" json:"clientParameters" yaml:"clientParameters"`
}

