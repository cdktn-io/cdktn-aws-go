package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallFirewallPolicy_RuleVariablesProperty struct {
	// ip_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#ip_set AwsNetworkfirewallFirewallPolicy#ip_set}
	// Experimental.
	IpSet *AwsNetworkfirewallFirewallPolicy_IpSetProperty `field:"required" json:"ipSet" yaml:"ipSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#key AwsNetworkfirewallFirewallPolicy#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

