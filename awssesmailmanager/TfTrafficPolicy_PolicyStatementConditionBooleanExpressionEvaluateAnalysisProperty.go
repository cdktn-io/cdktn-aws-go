package awssesmailmanager


// Experimental.
type TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#analyzer TfTrafficPolicy#analyzer}.
	// Experimental.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#result_field TfTrafficPolicy#result_field}.
	// Experimental.
	ResultField *string `field:"required" json:"resultField" yaml:"resultField"`
}

