package awsbedrockagentcore


// Experimental.
type TfGatewayRule_RouteToTargetProperty struct {
	// static_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#static_route TfGatewayRule#static_route}
	// Experimental.
	StaticRoute interface{} `field:"optional" json:"staticRoute" yaml:"staticRoute"`
	// weighted_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#weighted_route TfGatewayRule#weighted_route}
	// Experimental.
	WeightedRoute interface{} `field:"optional" json:"weightedRoute" yaml:"weightedRoute"`
}

