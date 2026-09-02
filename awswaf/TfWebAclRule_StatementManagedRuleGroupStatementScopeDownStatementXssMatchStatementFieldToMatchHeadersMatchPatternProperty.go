package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeadersMatchPatternProperty struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#all TfWebAclRule#all}
	// Experimental.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#excluded_headers TfWebAclRule#excluded_headers}.
	// Experimental.
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#included_headers TfWebAclRule#included_headers}.
	// Experimental.
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

