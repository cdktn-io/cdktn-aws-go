package awsnetworkfirewall


// Experimental.
type TfRuleGroup_RuleDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#actions TfRuleGroup#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// match_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#match_attributes TfRuleGroup#match_attributes}
	// Experimental.
	MatchAttributes *TfRuleGroup_MatchAttributesProperty `field:"required" json:"matchAttributes" yaml:"matchAttributes"`
}

