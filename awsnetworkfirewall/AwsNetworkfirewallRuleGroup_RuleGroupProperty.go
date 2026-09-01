package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_RuleGroupProperty struct {
	// rules_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rules_source AwsNetworkfirewallRuleGroup#rules_source}
	// Experimental.
	RulesSource *AwsNetworkfirewallRuleGroup_RulesSourceProperty `field:"required" json:"rulesSource" yaml:"rulesSource"`
	// reference_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#reference_sets AwsNetworkfirewallRuleGroup#reference_sets}
	// Experimental.
	ReferenceSets *AwsNetworkfirewallRuleGroup_ReferenceSetsProperty `field:"optional" json:"referenceSets" yaml:"referenceSets"`
	// rule_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rule_variables AwsNetworkfirewallRuleGroup#rule_variables}
	// Experimental.
	RuleVariables *AwsNetworkfirewallRuleGroup_RuleVariablesProperty `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// stateful_rule_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateful_rule_options AwsNetworkfirewallRuleGroup#stateful_rule_options}
	// Experimental.
	StatefulRuleOptions *AwsNetworkfirewallRuleGroup_StatefulRuleOptionsProperty `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}

