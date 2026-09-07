package apigatewayv2


// Experimental.
type AwsRoutingRule_ConditionProperty struct {
	// match_base_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#match_base_paths AwsRoutingRule#match_base_paths}
	// Experimental.
	MatchBasePaths interface{} `field:"optional" json:"matchBasePaths" yaml:"matchBasePaths"`
	// match_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#match_headers AwsRoutingRule#match_headers}
	// Experimental.
	MatchHeaders interface{} `field:"optional" json:"matchHeaders" yaml:"matchHeaders"`
}

