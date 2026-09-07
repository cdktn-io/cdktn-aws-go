package sesmailmanager


// Experimental.
type AwsRuleSet_RuleConditionStringExpressionEvaluateAnalysisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#analyzer AwsRuleSet#analyzer}.
	// Experimental.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#result_field AwsRuleSet#result_field}.
	// Experimental.
	ResultField *string `field:"required" json:"resultField" yaml:"resultField"`
}

