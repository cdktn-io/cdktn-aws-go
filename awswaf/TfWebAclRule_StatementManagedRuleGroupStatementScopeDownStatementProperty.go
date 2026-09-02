package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementProperty struct {
	// asn_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#asn_match_statement TfWebAclRule#asn_match_statement}
	// Experimental.
	AsnMatchStatement interface{} `field:"optional" json:"asnMatchStatement" yaml:"asnMatchStatement"`
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#byte_match_statement TfWebAclRule#byte_match_statement}
	// Experimental.
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#geo_match_statement TfWebAclRule#geo_match_statement}
	// Experimental.
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ip_set_reference_statement TfWebAclRule#ip_set_reference_statement}
	// Experimental.
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#label_match_statement TfWebAclRule#label_match_statement}
	// Experimental.
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#regex_match_statement TfWebAclRule#regex_match_statement}
	// Experimental.
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#regex_pattern_set_reference_statement TfWebAclRule#regex_pattern_set_reference_statement}
	// Experimental.
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#size_constraint_statement TfWebAclRule#size_constraint_statement}
	// Experimental.
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#sqli_match_statement TfWebAclRule#sqli_match_statement}
	// Experimental.
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#xss_match_statement TfWebAclRule#xss_match_statement}
	// Experimental.
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

