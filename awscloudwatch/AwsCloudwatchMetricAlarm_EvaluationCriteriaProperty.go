package awscloudwatch


// Experimental.
type AwsCloudwatchMetricAlarm_EvaluationCriteriaProperty struct {
	// promql_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#promql_criteria AwsCloudwatchMetricAlarm#promql_criteria}
	// Experimental.
	PromqlCriteria *AwsCloudwatchMetricAlarm_PromqlCriteriaProperty `field:"required" json:"promqlCriteria" yaml:"promqlCriteria"`
}

