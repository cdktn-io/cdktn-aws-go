package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallFirewallPolicy_StatefulRuleGroupReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#resource_arn AwsNetworkfirewallFirewallPolicy#resource_arn}.
	// Experimental.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#deep_threat_inspection AwsNetworkfirewallFirewallPolicy#deep_threat_inspection}.
	// Experimental.
	DeepThreatInspection *string `field:"optional" json:"deepThreatInspection" yaml:"deepThreatInspection"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#override AwsNetworkfirewallFirewallPolicy#override}
	// Experimental.
	Override *AwsNetworkfirewallFirewallPolicy_OverrideProperty `field:"optional" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#priority AwsNetworkfirewallFirewallPolicy#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

