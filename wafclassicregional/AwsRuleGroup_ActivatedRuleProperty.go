package wafclassicregional


// Experimental.
type AwsRuleGroup_ActivatedRuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#action AwsRuleGroup#action}
	// Experimental.
	Action *AwsRuleGroup_ActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#priority AwsRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#rule_id AwsRuleGroup#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#type AwsRuleGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

