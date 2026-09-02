package awsnetworkfirewall


// Experimental.
type TfFirewallPolicy_RuleVariablesProperty struct {
	// ip_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#ip_set TfFirewallPolicy#ip_set}
	// Experimental.
	IpSet *TfFirewallPolicy_IpSetProperty `field:"required" json:"ipSet" yaml:"ipSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#key TfFirewallPolicy#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

