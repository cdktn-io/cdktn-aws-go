package awssesmailmanager


// Experimental.
type TfRuleSet_RuleUnlessBooleanExpressionEvaluateProperty struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#analysis TfRuleSet#analysis}
	// Experimental.
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#attribute TfRuleSet#attribute}.
	// Experimental.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// is_in_address_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#is_in_address_list TfRuleSet#is_in_address_list}
	// Experimental.
	IsInAddressList interface{} `field:"optional" json:"isInAddressList" yaml:"isInAddressList"`
}

