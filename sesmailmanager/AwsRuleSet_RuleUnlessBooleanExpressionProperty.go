package sesmailmanager


// Experimental.
type AwsRuleSet_RuleUnlessBooleanExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#operator AwsRuleSet#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// evaluate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#evaluate AwsRuleSet#evaluate}
	// Experimental.
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
}

