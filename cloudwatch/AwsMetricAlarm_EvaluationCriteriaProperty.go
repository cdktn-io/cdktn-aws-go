package cloudwatch


// Experimental.
type AwsMetricAlarm_EvaluationCriteriaProperty struct {
	// promql_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#promql_criteria AwsMetricAlarm#promql_criteria}
	// Experimental.
	PromqlCriteria *AwsMetricAlarm_PromqlCriteriaProperty `field:"required" json:"promqlCriteria" yaml:"promqlCriteria"`
}

