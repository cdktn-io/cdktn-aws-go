package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_IpSetsProperty struct {
	// ip_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#ip_set AwsNetworkfirewallRuleGroup#ip_set}
	// Experimental.
	IpSet *AwsNetworkfirewallRuleGroup_IpSetProperty `field:"required" json:"ipSet" yaml:"ipSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#key AwsNetworkfirewallRuleGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

