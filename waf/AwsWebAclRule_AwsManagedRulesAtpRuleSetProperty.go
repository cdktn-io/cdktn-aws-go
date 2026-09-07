package waf


// Experimental.
type AwsWebAclRule_AwsManagedRulesAtpRuleSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#login_path AwsWebAclRule#login_path}.
	// Experimental.
	LoginPath *string `field:"required" json:"loginPath" yaml:"loginPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#enable_regex_in_path AwsWebAclRule#enable_regex_in_path}.
	// Experimental.
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// request_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#request_inspection AwsWebAclRule#request_inspection}
	// Experimental.
	RequestInspection interface{} `field:"optional" json:"requestInspection" yaml:"requestInspection"`
	// response_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#response_inspection AwsWebAclRule#response_inspection}
	// Experimental.
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

