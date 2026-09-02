package awsnetworkfirewall


// Experimental.
type TfRuleGroup_PortSetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#key TfRuleGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// port_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#port_set TfRuleGroup#port_set}
	// Experimental.
	PortSet *TfRuleGroup_PortSetProperty `field:"required" json:"portSet" yaml:"portSet"`
}

