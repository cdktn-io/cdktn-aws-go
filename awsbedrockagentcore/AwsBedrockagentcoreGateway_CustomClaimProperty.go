package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGateway_CustomClaimProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#inbound_token_claim_name AwsBedrockagentcoreGateway#inbound_token_claim_name}.
	// Experimental.
	InboundTokenClaimName *string `field:"required" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#inbound_token_claim_value_type AwsBedrockagentcoreGateway#inbound_token_claim_value_type}.
	// Experimental.
	InboundTokenClaimValueType *string `field:"required" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
	// authorizing_claim_match_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#authorizing_claim_match_value AwsBedrockagentcoreGateway#authorizing_claim_match_value}
	// Experimental.
	AuthorizingClaimMatchValue interface{} `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
}

