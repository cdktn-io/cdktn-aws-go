package networkfirewall


// Experimental.
type AwsRuleGroup_RuleDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#actions AwsRuleGroup#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// match_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#match_attributes AwsRuleGroup#match_attributes}
	// Experimental.
	MatchAttributes *AwsRuleGroup_MatchAttributesProperty `field:"required" json:"matchAttributes" yaml:"matchAttributes"`
}

