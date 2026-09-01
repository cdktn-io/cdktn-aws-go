package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreAgentRuntime_CustomClaimProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#inbound_token_claim_name AwsBedrockagentcoreAgentRuntime#inbound_token_claim_name}.
	// Experimental.
	InboundTokenClaimName *string `field:"required" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#inbound_token_claim_value_type AwsBedrockagentcoreAgentRuntime#inbound_token_claim_value_type}.
	// Experimental.
	InboundTokenClaimValueType *string `field:"required" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
	// authorizing_claim_match_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#authorizing_claim_match_value AwsBedrockagentcoreAgentRuntime#authorizing_claim_match_value}
	// Experimental.
	AuthorizingClaimMatchValue interface{} `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
}

