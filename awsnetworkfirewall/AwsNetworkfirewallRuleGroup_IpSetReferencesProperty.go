package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_IpSetReferencesProperty struct {
	// ip_set_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#ip_set_reference AwsNetworkfirewallRuleGroup#ip_set_reference}
	// Experimental.
	IpSetReference interface{} `field:"required" json:"ipSetReference" yaml:"ipSetReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#key AwsNetworkfirewallRuleGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

