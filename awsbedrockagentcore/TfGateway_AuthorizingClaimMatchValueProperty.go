package awsbedrockagentcore


// Experimental.
type TfGateway_AuthorizingClaimMatchValueProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#claim_match_operator TfGateway#claim_match_operator}.
	// Experimental.
	ClaimMatchOperator *string `field:"required" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// claim_match_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#claim_match_value TfGateway#claim_match_value}
	// Experimental.
	ClaimMatchValue interface{} `field:"optional" json:"claimMatchValue" yaml:"claimMatchValue"`
}

