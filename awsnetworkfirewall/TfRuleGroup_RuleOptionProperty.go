package awsnetworkfirewall


// Experimental.
type TfRuleGroup_RuleOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#keyword TfRuleGroup#keyword}.
	// Experimental.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#settings TfRuleGroup#settings}.
	// Experimental.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

