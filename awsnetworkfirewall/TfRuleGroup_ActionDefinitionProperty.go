package awsnetworkfirewall


// Experimental.
type TfRuleGroup_ActionDefinitionProperty struct {
	// publish_metric_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#publish_metric_action TfRuleGroup#publish_metric_action}
	// Experimental.
	PublishMetricAction *TfRuleGroup_PublishMetricActionProperty `field:"required" json:"publishMetricAction" yaml:"publishMetricAction"`
}

