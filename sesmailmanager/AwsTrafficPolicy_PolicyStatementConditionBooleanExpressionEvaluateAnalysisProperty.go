package sesmailmanager


// Experimental.
type AwsTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#analyzer AwsTrafficPolicy#analyzer}.
	// Experimental.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#result_field AwsTrafficPolicy#result_field}.
	// Experimental.
	ResultField *string `field:"required" json:"resultField" yaml:"resultField"`
}

