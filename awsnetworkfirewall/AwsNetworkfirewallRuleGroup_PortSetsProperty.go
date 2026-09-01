package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_PortSetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#key AwsNetworkfirewallRuleGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// port_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#port_set AwsNetworkfirewallRuleGroup#port_set}
	// Experimental.
	PortSet *AwsNetworkfirewallRuleGroup_PortSetProperty `field:"required" json:"portSet" yaml:"portSet"`
}

