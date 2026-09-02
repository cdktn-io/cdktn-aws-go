package awsbedrockagentcore


// Experimental.
type TfRegistry_PrivateEndpointOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#domain TfRegistry#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#private_endpoint TfRegistry#private_endpoint}
	// Experimental.
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

