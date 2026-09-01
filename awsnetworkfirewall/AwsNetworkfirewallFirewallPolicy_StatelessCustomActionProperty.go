package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallFirewallPolicy_StatelessCustomActionProperty struct {
	// action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#action_definition AwsNetworkfirewallFirewallPolicy#action_definition}
	// Experimental.
	ActionDefinition *AwsNetworkfirewallFirewallPolicy_ActionDefinitionProperty `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#action_name AwsNetworkfirewallFirewallPolicy#action_name}.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}

