package awsnetworkfirewall


// Experimental.
type TfRuleGroup_StatefulRuleOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#rule_order TfRuleGroup#rule_order}.
	// Experimental.
	RuleOrder *string `field:"required" json:"ruleOrder" yaml:"ruleOrder"`
}

