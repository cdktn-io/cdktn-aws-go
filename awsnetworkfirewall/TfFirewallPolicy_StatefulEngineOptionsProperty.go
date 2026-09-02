package awsnetworkfirewall


// Experimental.
type TfFirewallPolicy_StatefulEngineOptionsProperty struct {
	// flow_timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#flow_timeouts TfFirewallPolicy#flow_timeouts}
	// Experimental.
	FlowTimeouts *TfFirewallPolicy_FlowTimeoutsProperty `field:"optional" json:"flowTimeouts" yaml:"flowTimeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#rule_order TfFirewallPolicy#rule_order}.
	// Experimental.
	RuleOrder *string `field:"optional" json:"ruleOrder" yaml:"ruleOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stream_exception_policy TfFirewallPolicy#stream_exception_policy}.
	// Experimental.
	StreamExceptionPolicy *string `field:"optional" json:"streamExceptionPolicy" yaml:"streamExceptionPolicy"`
}

