package waf


// Experimental.
type AwsWebAclLoggingConfiguration_RedactedFieldsProperty struct {
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#method AwsWebAclLoggingConfiguration#method}
	// Experimental.
	Method *AwsWebAclLoggingConfiguration_MethodProperty `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#query_string AwsWebAclLoggingConfiguration#query_string}
	// Experimental.
	QueryString *AwsWebAclLoggingConfiguration_QueryStringProperty `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#single_header AwsWebAclLoggingConfiguration#single_header}
	// Experimental.
	SingleHeader *AwsWebAclLoggingConfiguration_SingleHeaderProperty `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#uri_path AwsWebAclLoggingConfiguration#uri_path}
	// Experimental.
	UriPath *AwsWebAclLoggingConfiguration_UriPathProperty `field:"optional" json:"uriPath" yaml:"uriPath"`
}

