package awswaf


// Experimental.
type TfWebAclRule_StatementByteMatchStatementProperty struct {
	// Area within the portion of a web request that you want AWS WAF to search for SearchString.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#positional_constraint TfWebAclRule#positional_constraint}
	// Experimental.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// String value to search for within the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#search_string TfWebAclRule#search_string}
	// Experimental.
	SearchString *string `field:"required" json:"searchString" yaml:"searchString"`
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

