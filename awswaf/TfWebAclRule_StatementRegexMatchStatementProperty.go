package awswaf


// Experimental.
type TfWebAclRule_StatementRegexMatchStatementProperty struct {
	// Regular expression string (1-512 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#regex_string TfWebAclRule#regex_string}
	// Experimental.
	RegexString *string `field:"required" json:"regexString" yaml:"regexString"`
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

