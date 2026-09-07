package waf


// Experimental.
type AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementProperty struct {
	// ARN of the RegexPatternSet (20-2048 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#arn AwsWebAclRule#arn}
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#field_to_match AwsWebAclRule#field_to_match}
	// Experimental.
	FieldToMatch interface{} `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#text_transformation AwsWebAclRule#text_transformation}
	// Experimental.
	TextTransformation interface{} `field:"optional" json:"textTransformation" yaml:"textTransformation"`
}

