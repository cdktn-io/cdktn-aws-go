package awsnetworkfirewall


// Experimental.
type TfRuleGroup_RuleVariablesProperty struct {
	// ip_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#ip_sets TfRuleGroup#ip_sets}
	// Experimental.
	IpSets interface{} `field:"optional" json:"ipSets" yaml:"ipSets"`
	// port_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#port_sets TfRuleGroup#port_sets}
	// Experimental.
	PortSets interface{} `field:"optional" json:"portSets" yaml:"portSets"`
}

