package awswaf


// Experimental.
type TfWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#field_to_match TfWebAclRule#field_to_match}
	// Experimental.
	FieldToMatch interface{} `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#text_transformation TfWebAclRule#text_transformation}
	// Experimental.
	TextTransformation interface{} `field:"optional" json:"textTransformation" yaml:"textTransformation"`
}

