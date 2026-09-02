package awsnetworkfirewall


// Experimental.
type TfFirewallPolicy_ActionDefinitionProperty struct {
	// publish_metric_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#publish_metric_action TfFirewallPolicy#publish_metric_action}
	// Experimental.
	PublishMetricAction *TfFirewallPolicy_PublishMetricActionProperty `field:"required" json:"publishMetricAction" yaml:"publishMetricAction"`
}

