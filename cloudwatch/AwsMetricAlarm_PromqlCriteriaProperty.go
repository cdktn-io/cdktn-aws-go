package cloudwatch


// Experimental.
type AwsMetricAlarm_PromqlCriteriaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#query AwsMetricAlarm#query}.
	// Experimental.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#pending_period AwsMetricAlarm#pending_period}.
	// Experimental.
	PendingPeriod *float64 `field:"optional" json:"pendingPeriod" yaml:"pendingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#recovery_period AwsMetricAlarm#recovery_period}.
	// Experimental.
	RecoveryPeriod *float64 `field:"optional" json:"recoveryPeriod" yaml:"recoveryPeriod"`
}

