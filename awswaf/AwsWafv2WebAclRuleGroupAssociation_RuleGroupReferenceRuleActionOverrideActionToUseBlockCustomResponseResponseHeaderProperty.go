package awswaf


// Experimental.
type AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#name AwsWafv2WebAclRuleGroupAssociation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#value AwsWafv2WebAclRuleGroupAssociation#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

