package networkfirewall


// Experimental.
type AwsRuleGroup_RuleOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#keyword AwsRuleGroup#keyword}.
	// Experimental.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#settings AwsRuleGroup#settings}.
	// Experimental.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

