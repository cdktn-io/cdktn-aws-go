package awsnetworkfirewall


// Experimental.
type TfRuleGroup_RuleGroupProperty struct {
	// rules_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rules_source TfRuleGroup#rules_source}
	// Experimental.
	RulesSource *TfRuleGroup_RulesSourceProperty `field:"required" json:"rulesSource" yaml:"rulesSource"`
	// reference_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#reference_sets TfRuleGroup#reference_sets}
	// Experimental.
	ReferenceSets *TfRuleGroup_ReferenceSetsProperty `field:"optional" json:"referenceSets" yaml:"referenceSets"`
	// rule_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rule_variables TfRuleGroup#rule_variables}
	// Experimental.
	RuleVariables *TfRuleGroup_RuleVariablesProperty `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// stateful_rule_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateful_rule_options TfRuleGroup#stateful_rule_options}
	// Experimental.
	StatefulRuleOptions *TfRuleGroup_StatefulRuleOptionsProperty `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}

