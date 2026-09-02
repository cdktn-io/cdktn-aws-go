package awsnetworkfirewall


// Experimental.
type TfRuleGroup_StatelessRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#priority TfRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rule_definition TfRuleGroup#rule_definition}
	// Experimental.
	RuleDefinition *TfRuleGroup_RuleDefinitionProperty `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
}

