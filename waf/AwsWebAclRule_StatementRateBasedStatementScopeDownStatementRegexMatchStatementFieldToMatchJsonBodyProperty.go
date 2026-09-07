package waf


// Experimental.
type AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#match_scope AwsWebAclRule#match_scope}.
	// Experimental.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#invalid_fallback_behavior AwsWebAclRule#invalid_fallback_behavior}.
	// Experimental.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
	// match_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#match_pattern AwsWebAclRule#match_pattern}
	// Experimental.
	MatchPattern interface{} `field:"optional" json:"matchPattern" yaml:"matchPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#oversize_handling AwsWebAclRule#oversize_handling}.
	// Experimental.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

