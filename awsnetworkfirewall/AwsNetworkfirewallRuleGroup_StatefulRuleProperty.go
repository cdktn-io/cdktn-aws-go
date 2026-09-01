package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallRuleGroup_StatefulRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#action AwsNetworkfirewallRuleGroup#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#header AwsNetworkfirewallRuleGroup#header}
	// Experimental.
	Header *AwsNetworkfirewallRuleGroup_HeaderProperty `field:"required" json:"header" yaml:"header"`
	// rule_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rule_option AwsNetworkfirewallRuleGroup#rule_option}
	// Experimental.
	RuleOption interface{} `field:"required" json:"ruleOption" yaml:"ruleOption"`
}

