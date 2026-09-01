package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayRule_RouteToTargetProperty struct {
	// static_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#static_route AwsBedrockagentcoreGatewayRule#static_route}
	// Experimental.
	StaticRoute interface{} `field:"optional" json:"staticRoute" yaml:"staticRoute"`
	// weighted_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#weighted_route AwsBedrockagentcoreGatewayRule#weighted_route}
	// Experimental.
	WeightedRoute interface{} `field:"optional" json:"weightedRoute" yaml:"weightedRoute"`
}

