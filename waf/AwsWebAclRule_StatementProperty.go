package waf


// Experimental.
type AwsWebAclRule_StatementProperty struct {
	// and_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#and_statement AwsWebAclRule#and_statement}
	// Experimental.
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// asn_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#asn_match_statement AwsWebAclRule#asn_match_statement}
	// Experimental.
	AsnMatchStatement interface{} `field:"optional" json:"asnMatchStatement" yaml:"asnMatchStatement"`
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#byte_match_statement AwsWebAclRule#byte_match_statement}
	// Experimental.
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#geo_match_statement AwsWebAclRule#geo_match_statement}
	// Experimental.
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ip_set_reference_statement AwsWebAclRule#ip_set_reference_statement}
	// Experimental.
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#label_match_statement AwsWebAclRule#label_match_statement}
	// Experimental.
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// managed_rule_group_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#managed_rule_group_statement AwsWebAclRule#managed_rule_group_statement}
	// Experimental.
	ManagedRuleGroupStatement interface{} `field:"optional" json:"managedRuleGroupStatement" yaml:"managedRuleGroupStatement"`
	// not_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#not_statement AwsWebAclRule#not_statement}
	// Experimental.
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// or_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#or_statement AwsWebAclRule#or_statement}
	// Experimental.
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
	// rate_based_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#rate_based_statement AwsWebAclRule#rate_based_statement}
	// Experimental.
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#regex_match_statement AwsWebAclRule#regex_match_statement}
	// Experimental.
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#regex_pattern_set_reference_statement AwsWebAclRule#regex_pattern_set_reference_statement}
	// Experimental.
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// rule_group_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#rule_group_reference_statement AwsWebAclRule#rule_group_reference_statement}
	// Experimental.
	RuleGroupReferenceStatement interface{} `field:"optional" json:"ruleGroupReferenceStatement" yaml:"ruleGroupReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#size_constraint_statement AwsWebAclRule#size_constraint_statement}
	// Experimental.
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#sqli_match_statement AwsWebAclRule#sqli_match_statement}
	// Experimental.
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#xss_match_statement AwsWebAclRule#xss_match_statement}
	// Experimental.
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

