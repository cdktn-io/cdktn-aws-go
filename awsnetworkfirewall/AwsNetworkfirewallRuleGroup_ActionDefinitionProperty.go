package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_ActionDefinitionProperty struct {
	// publish_metric_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#publish_metric_action AwsNetworkfirewallRuleGroup#publish_metric_action}
	// Experimental.
	PublishMetricAction *AwsNetworkfirewallRuleGroup_PublishMetricActionProperty `field:"required" json:"publishMetricAction" yaml:"publishMetricAction"`
}

