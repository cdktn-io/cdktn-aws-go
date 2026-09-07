package networkfirewall


// Experimental.
type AwsFirewallPolicy_StatelessCustomActionProperty struct {
	// action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#action_definition AwsFirewallPolicy#action_definition}
	// Experimental.
	ActionDefinition *AwsFirewallPolicy_ActionDefinitionProperty `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#action_name AwsFirewallPolicy#action_name}.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}

