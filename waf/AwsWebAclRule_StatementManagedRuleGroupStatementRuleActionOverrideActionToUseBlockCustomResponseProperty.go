package waf


// Experimental.
type AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlockCustomResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#response_code AwsWebAclRule#response_code}.
	// Experimental.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#custom_response_body_key AwsWebAclRule#custom_response_body_key}.
	// Experimental.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#response_header AwsWebAclRule#response_header}
	// Experimental.
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}

