package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_StatelessRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#priority AwsNetworkfirewallRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rule_definition AwsNetworkfirewallRuleGroup#rule_definition}
	// Experimental.
	RuleDefinition *AwsNetworkfirewallRuleGroup_RuleDefinitionProperty `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
}

