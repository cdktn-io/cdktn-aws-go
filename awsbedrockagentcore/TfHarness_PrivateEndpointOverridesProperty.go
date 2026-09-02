package awsbedrockagentcore


// Experimental.
type TfHarness_PrivateEndpointOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#domain TfHarness#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#private_endpoint TfHarness#private_endpoint}
	// Experimental.
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

