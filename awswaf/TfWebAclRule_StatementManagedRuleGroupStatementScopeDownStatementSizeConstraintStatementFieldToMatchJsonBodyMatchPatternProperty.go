package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementFieldToMatchJsonBodyMatchPatternProperty struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#all TfWebAclRule#all}
	// Experimental.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#included_paths TfWebAclRule#included_paths}.
	// Experimental.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

