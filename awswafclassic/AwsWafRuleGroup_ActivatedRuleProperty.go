package awswafclassic


// Experimental.
type AwsWafRuleGroup_ActivatedRuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#action AwsWafRuleGroup#action}
	// Experimental.
	Action *AwsWafRuleGroup_ActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#priority AwsWafRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#rule_id AwsWafRuleGroup#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#type AwsWafRuleGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

