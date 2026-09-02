package awswaf


// Experimental.
type TfWebAclLoggingConfiguration_RedactedFieldsProperty struct {
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#method TfWebAclLoggingConfiguration#method}
	// Experimental.
	Method *TfWebAclLoggingConfiguration_MethodProperty `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#query_string TfWebAclLoggingConfiguration#query_string}
	// Experimental.
	QueryString *TfWebAclLoggingConfiguration_QueryStringProperty `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#single_header TfWebAclLoggingConfiguration#single_header}
	// Experimental.
	SingleHeader *TfWebAclLoggingConfiguration_SingleHeaderProperty `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#uri_path TfWebAclLoggingConfiguration#uri_path}
	// Experimental.
	UriPath *TfWebAclLoggingConfiguration_UriPathProperty `field:"optional" json:"uriPath" yaml:"uriPath"`
}

