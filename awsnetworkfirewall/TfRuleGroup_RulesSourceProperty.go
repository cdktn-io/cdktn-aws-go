package awsnetworkfirewall


// Experimental.
type TfRuleGroup_RulesSourceProperty struct {
	// rules_source_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rules_source_list TfRuleGroup#rules_source_list}
	// Experimental.
	RulesSourceList *TfRuleGroup_RulesSourceListProperty `field:"optional" json:"rulesSourceList" yaml:"rulesSourceList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rules_string TfRuleGroup#rules_string}.
	// Experimental.
	RulesString *string `field:"optional" json:"rulesString" yaml:"rulesString"`
	// stateful_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateful_rule TfRuleGroup#stateful_rule}
	// Experimental.
	StatefulRule interface{} `field:"optional" json:"statefulRule" yaml:"statefulRule"`
	// stateless_rules_and_custom_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateless_rules_and_custom_actions TfRuleGroup#stateless_rules_and_custom_actions}
	// Experimental.
	StatelessRulesAndCustomActions *TfRuleGroup_StatelessRulesAndCustomActionsProperty `field:"optional" json:"statelessRulesAndCustomActions" yaml:"statelessRulesAndCustomActions"`
}

