package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_CustomClaimProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#inbound_token_claim_name AwsBedrockagentcoreHarness#inbound_token_claim_name}.
	// Experimental.
	InboundTokenClaimName *string `field:"required" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#inbound_token_claim_value_type AwsBedrockagentcoreHarness#inbound_token_claim_value_type}.
	// Experimental.
	InboundTokenClaimValueType *string `field:"required" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
	// authorizing_claim_match_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#authorizing_claim_match_value AwsBedrockagentcoreHarness#authorizing_claim_match_value}
	// Experimental.
	AuthorizingClaimMatchValue interface{} `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
}

