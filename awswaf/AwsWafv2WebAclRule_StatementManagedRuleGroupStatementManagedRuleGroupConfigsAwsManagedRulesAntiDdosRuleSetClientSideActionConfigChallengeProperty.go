package awswaf


// Experimental.
type AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAntiDdosRuleSetClientSideActionConfigChallengeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#usage_of_action AwsWafv2WebAclRule#usage_of_action}.
	// Experimental.
	UsageOfAction *string `field:"required" json:"usageOfAction" yaml:"usageOfAction"`
	// exempt_uri_regular_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#exempt_uri_regular_expression AwsWafv2WebAclRule#exempt_uri_regular_expression}
	// Experimental.
	ExemptUriRegularExpression interface{} `field:"optional" json:"exemptUriRegularExpression" yaml:"exemptUriRegularExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#sensitivity AwsWafv2WebAclRule#sensitivity}.
	// Experimental.
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
}

