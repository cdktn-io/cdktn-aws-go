package awswaf


// Experimental.
type AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementProperty struct {
	// List of two-character country codes (e.g., US, CA).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#country_codes AwsWafv2WebAclRule#country_codes}
	// Experimental.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#forwarded_ip_config AwsWafv2WebAclRule#forwarded_ip_config}
	// Experimental.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

