package awssesmailmanager


// Experimental.
type TfRuleSet_RuleUnlessVerdictExpressionEvaluateProperty struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#analysis TfRuleSet#analysis}
	// Experimental.
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#attribute TfRuleSet#attribute}.
	// Experimental.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

