package awsbedrockagentcore


// Experimental.
type TfGatewayRule_ActionConfigurationBundleWeightedOverrideTrafficSplitProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#name TfGatewayRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#weight TfGatewayRule#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// configuration_bundle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#configuration_bundle TfGatewayRule#configuration_bundle}
	// Experimental.
	ConfigurationBundle interface{} `field:"optional" json:"configurationBundle" yaml:"configurationBundle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#description TfGatewayRule#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#metadata TfGatewayRule#metadata}.
	// Experimental.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

