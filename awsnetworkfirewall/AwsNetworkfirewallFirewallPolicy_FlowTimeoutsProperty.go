package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallFirewallPolicy_FlowTimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#tcp_idle_timeout_seconds AwsNetworkfirewallFirewallPolicy#tcp_idle_timeout_seconds}.
	// Experimental.
	TcpIdleTimeoutSeconds *float64 `field:"optional" json:"tcpIdleTimeoutSeconds" yaml:"tcpIdleTimeoutSeconds"`
}

