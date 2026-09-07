package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAntiDdosRuleSetClientSideActionConfigChallengeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#usage_of_action AwsWebAclRuleGroupAssociation#usage_of_action}.
	// Experimental.
	UsageOfAction *string `field:"required" json:"usageOfAction" yaml:"usageOfAction"`
	// exempt_uri_regular_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#exempt_uri_regular_expression AwsWebAclRuleGroupAssociation#exempt_uri_regular_expression}
	// Experimental.
	ExemptUriRegularExpression interface{} `field:"optional" json:"exemptUriRegularExpression" yaml:"exemptUriRegularExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#sensitivity AwsWebAclRuleGroupAssociation#sensitivity}.
	// Experimental.
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
}

