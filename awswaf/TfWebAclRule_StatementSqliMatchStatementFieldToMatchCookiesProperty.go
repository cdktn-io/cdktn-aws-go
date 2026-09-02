package awswaf


// Experimental.
type TfWebAclRule_StatementSqliMatchStatementFieldToMatchCookiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#match_scope TfWebAclRule#match_scope}.
	// Experimental.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#oversize_handling TfWebAclRule#oversize_handling}.
	// Experimental.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
	// match_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#match_pattern TfWebAclRule#match_pattern}
	// Experimental.
	MatchPattern interface{} `field:"optional" json:"matchPattern" yaml:"matchPattern"`
}

