package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingPolicy_PredictiveScalingPolicyConfigurationProperty struct {
	// metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_specification AwsAppautoscalingPolicy#metric_specification}
	// Experimental.
	MetricSpecification interface{} `field:"required" json:"metricSpecification" yaml:"metricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#max_capacity_breach_behavior AwsAppautoscalingPolicy#max_capacity_breach_behavior}.
	// Experimental.
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#max_capacity_buffer AwsAppautoscalingPolicy#max_capacity_buffer}.
	// Experimental.
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#mode AwsAppautoscalingPolicy#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scheduling_buffer_time AwsAppautoscalingPolicy#scheduling_buffer_time}.
	// Experimental.
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

