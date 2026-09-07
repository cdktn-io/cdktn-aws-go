package waf


// Experimental.
type AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#priority AwsWebAclRule#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#type AwsWebAclRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

