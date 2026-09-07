package networkfirewall


// Experimental.
type AwsRuleGroup_RulesSourceProperty struct {
	// rules_source_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rules_source_list AwsRuleGroup#rules_source_list}
	// Experimental.
	RulesSourceList *AwsRuleGroup_RulesSourceListProperty `field:"optional" json:"rulesSourceList" yaml:"rulesSourceList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rules_string AwsRuleGroup#rules_string}.
	// Experimental.
	RulesString *string `field:"optional" json:"rulesString" yaml:"rulesString"`
	// stateful_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateful_rule AwsRuleGroup#stateful_rule}
	// Experimental.
	StatefulRule interface{} `field:"optional" json:"statefulRule" yaml:"statefulRule"`
	// stateless_rules_and_custom_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateless_rules_and_custom_actions AwsRuleGroup#stateless_rules_and_custom_actions}
	// Experimental.
	StatelessRulesAndCustomActions *AwsRuleGroup_StatelessRulesAndCustomActionsProperty `field:"optional" json:"statelessRulesAndCustomActions" yaml:"statelessRulesAndCustomActions"`
}

