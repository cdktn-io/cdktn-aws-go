package awsbedrockagentcore


// Experimental.
type TfGateway_PrivateEndpointOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#domain TfGateway#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#private_endpoint TfGateway#private_endpoint}
	// Experimental.
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

