package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingPolicy_StepScalingPolicyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#adjustment_type AwsAppautoscalingPolicy#adjustment_type}.
	// Experimental.
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#cooldown AwsAppautoscalingPolicy#cooldown}.
	// Experimental.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_aggregation_type AwsAppautoscalingPolicy#metric_aggregation_type}.
	// Experimental.
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#min_adjustment_magnitude AwsAppautoscalingPolicy#min_adjustment_magnitude}.
	// Experimental.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// step_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#step_adjustment AwsAppautoscalingPolicy#step_adjustment}
	// Experimental.
	StepAdjustment interface{} `field:"optional" json:"stepAdjustment" yaml:"stepAdjustment"`
}

