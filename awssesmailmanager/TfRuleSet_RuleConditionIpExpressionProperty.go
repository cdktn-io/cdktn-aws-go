package awssesmailmanager


// Experimental.
type TfRuleSet_RuleConditionIpExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#operator TfRuleSet#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#values TfRuleSet#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// evaluate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#evaluate TfRuleSet#evaluate}
	// Experimental.
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
}

