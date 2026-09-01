package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallFirewallPolicy_ActionDefinitionProperty struct {
	// publish_metric_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#publish_metric_action AwsNetworkfirewallFirewallPolicy#publish_metric_action}
	// Experimental.
	PublishMetricAction *AwsNetworkfirewallFirewallPolicy_PublishMetricActionProperty `field:"required" json:"publishMetricAction" yaml:"publishMetricAction"`
}

