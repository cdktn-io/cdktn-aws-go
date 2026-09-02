package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_MetadataConfigurationProperty struct {
	// A list of URL query parameters that are allowed to be propagated from incoming gateway URL to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#allowed_query_parameters TfGatewayTarget#allowed_query_parameters}
	// Experimental.
	AllowedQueryParameters *[]*string `field:"optional" json:"allowedQueryParameters" yaml:"allowedQueryParameters"`
	// A list of HTTP headers that are allowed to be propagated from incoming client requests to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#allowed_request_headers TfGatewayTarget#allowed_request_headers}
	// Experimental.
	AllowedRequestHeaders *[]*string `field:"optional" json:"allowedRequestHeaders" yaml:"allowedRequestHeaders"`
	// A list of HTTP headers that are allowed to be propagated from the target response back to the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#allowed_response_headers TfGatewayTarget#allowed_response_headers}
	// Experimental.
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
}

