package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideProperty struct {
	// Name of the rule to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#name AwsWebAclRuleGroupAssociation#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// action_to_use block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#action_to_use AwsWebAclRuleGroupAssociation#action_to_use}
	// Experimental.
	ActionToUse interface{} `field:"optional" json:"actionToUse" yaml:"actionToUse"`
}

