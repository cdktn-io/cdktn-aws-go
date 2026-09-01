package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_RuleOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#keyword AwsNetworkfirewallRuleGroup#keyword}.
	// Experimental.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#settings AwsNetworkfirewallRuleGroup#settings}.
	// Experimental.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

