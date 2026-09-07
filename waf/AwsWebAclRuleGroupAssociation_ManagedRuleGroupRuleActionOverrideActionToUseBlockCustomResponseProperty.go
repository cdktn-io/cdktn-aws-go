package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#response_code AwsWebAclRuleGroupAssociation#response_code}.
	// Experimental.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#custom_response_body_key AwsWebAclRuleGroupAssociation#custom_response_body_key}.
	// Experimental.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#response_header AwsWebAclRuleGroupAssociation#response_header}
	// Experimental.
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}

