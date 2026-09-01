package awswaf


// Experimental.
type AwsWafv2WebAclRule_RuleGroupReferenceStatementProperty struct {
	// ARN of the RuleGroup (20-2048 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#arn AwsWafv2WebAclRule#arn}
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// excluded_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#excluded_rule AwsWafv2WebAclRule#excluded_rule}
	// Experimental.
	ExcludedRule interface{} `field:"optional" json:"excludedRule" yaml:"excludedRule"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#rule_action_override AwsWafv2WebAclRule#rule_action_override}
	// Experimental.
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
}

