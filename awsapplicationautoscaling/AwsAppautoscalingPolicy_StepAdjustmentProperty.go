package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingPolicy_StepAdjustmentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scaling_adjustment AwsAppautoscalingPolicy#scaling_adjustment}.
	// Experimental.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_interval_lower_bound AwsAppautoscalingPolicy#metric_interval_lower_bound}.
	// Experimental.
	MetricIntervalLowerBound *string `field:"optional" json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_interval_upper_bound AwsAppautoscalingPolicy#metric_interval_upper_bound}.
	// Experimental.
	MetricIntervalUpperBound *string `field:"optional" json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

