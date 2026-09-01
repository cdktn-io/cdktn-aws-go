package awswafclassicregional


// Experimental.
type AwsWafregionalRuleGroup_ActivatedRuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#action AwsWafregionalRuleGroup#action}
	// Experimental.
	Action *AwsWafregionalRuleGroup_ActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#priority AwsWafregionalRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#rule_id AwsWafregionalRuleGroup#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule_group#type AwsWafregionalRuleGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

