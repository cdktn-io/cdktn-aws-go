package awswafclassic


// Experimental.
type TfRuleGroup_ActivatedRuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#action TfRuleGroup#action}
	// Experimental.
	Action *TfRuleGroup_ActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#priority TfRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#rule_id TfRuleGroup#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule_group#type TfRuleGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

