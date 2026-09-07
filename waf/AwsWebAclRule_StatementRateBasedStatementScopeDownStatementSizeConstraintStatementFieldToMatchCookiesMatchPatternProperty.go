package waf


// Experimental.
type AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementFieldToMatchCookiesMatchPatternProperty struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#all AwsWebAclRule#all}
	// Experimental.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#excluded_cookies AwsWebAclRule#excluded_cookies}.
	// Experimental.
	ExcludedCookies *[]*string `field:"optional" json:"excludedCookies" yaml:"excludedCookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#included_cookies AwsWebAclRule#included_cookies}.
	// Experimental.
	IncludedCookies *[]*string `field:"optional" json:"includedCookies" yaml:"includedCookies"`
}

