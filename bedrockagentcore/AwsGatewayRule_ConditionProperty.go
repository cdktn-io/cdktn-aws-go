package bedrockagentcore


// Experimental.
type AwsGatewayRule_ConditionProperty struct {
	// match_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#match_paths AwsGatewayRule#match_paths}
	// Experimental.
	MatchPaths interface{} `field:"optional" json:"matchPaths" yaml:"matchPaths"`
	// match_principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#match_principals AwsGatewayRule#match_principals}
	// Experimental.
	MatchPrincipals interface{} `field:"optional" json:"matchPrincipals" yaml:"matchPrincipals"`
}

