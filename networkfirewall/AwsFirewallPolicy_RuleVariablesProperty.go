package networkfirewall


// Experimental.
type AwsFirewallPolicy_RuleVariablesProperty struct {
	// ip_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#ip_set AwsFirewallPolicy#ip_set}
	// Experimental.
	IpSet *AwsFirewallPolicy_IpSetProperty `field:"required" json:"ipSet" yaml:"ipSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#key AwsFirewallPolicy#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

