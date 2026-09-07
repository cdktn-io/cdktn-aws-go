package waf


// Experimental.
type AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideProperty struct {
	// Name of the rule to override (1-128 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#name AwsWebAclRule#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// action_to_use block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#action_to_use AwsWebAclRule#action_to_use}
	// Experimental.
	ActionToUse interface{} `field:"optional" json:"actionToUse" yaml:"actionToUse"`
}

