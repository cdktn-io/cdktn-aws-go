package networkfirewall


// Experimental.
type AwsFirewallPolicy_FlowTimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#tcp_idle_timeout_seconds AwsFirewallPolicy#tcp_idle_timeout_seconds}.
	// Experimental.
	TcpIdleTimeoutSeconds *float64 `field:"optional" json:"tcpIdleTimeoutSeconds" yaml:"tcpIdleTimeoutSeconds"`
}

