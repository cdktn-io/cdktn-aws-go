package awswaf


// Experimental.
type AwsWafv2WebAclLoggingConfiguration_RedactedFieldsProperty struct {
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#method AwsWafv2WebAclLoggingConfiguration#method}
	// Experimental.
	Method *AwsWafv2WebAclLoggingConfiguration_MethodProperty `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#query_string AwsWafv2WebAclLoggingConfiguration#query_string}
	// Experimental.
	QueryString *AwsWafv2WebAclLoggingConfiguration_QueryStringProperty `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#single_header AwsWafv2WebAclLoggingConfiguration#single_header}
	// Experimental.
	SingleHeader *AwsWafv2WebAclLoggingConfiguration_SingleHeaderProperty `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#uri_path AwsWafv2WebAclLoggingConfiguration#uri_path}
	// Experimental.
	UriPath *AwsWafv2WebAclLoggingConfiguration_UriPathProperty `field:"optional" json:"uriPath" yaml:"uriPath"`
}

