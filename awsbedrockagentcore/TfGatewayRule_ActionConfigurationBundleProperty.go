package awsbedrockagentcore


// Experimental.
type TfGatewayRule_ActionConfigurationBundleProperty struct {
	// static_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#static_override TfGatewayRule#static_override}
	// Experimental.
	StaticOverride interface{} `field:"optional" json:"staticOverride" yaml:"staticOverride"`
	// weighted_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#weighted_override TfGatewayRule#weighted_override}
	// Experimental.
	WeightedOverride interface{} `field:"optional" json:"weightedOverride" yaml:"weightedOverride"`
}

