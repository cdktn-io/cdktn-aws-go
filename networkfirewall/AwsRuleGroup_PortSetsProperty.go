package networkfirewall


// Experimental.
type AwsRuleGroup_PortSetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#key AwsRuleGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// port_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#port_set AwsRuleGroup#port_set}
	// Experimental.
	PortSet *AwsRuleGroup_PortSetProperty `field:"required" json:"portSet" yaml:"portSet"`
}

