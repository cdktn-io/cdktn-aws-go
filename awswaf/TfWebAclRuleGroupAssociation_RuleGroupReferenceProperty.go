package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_RuleGroupReferenceProperty struct {
	// ARN of the Rule Group to associate with the Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#arn TfWebAclRuleGroupAssociation#arn}
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#rule_action_override TfWebAclRuleGroupAssociation#rule_action_override}
	// Experimental.
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
}

