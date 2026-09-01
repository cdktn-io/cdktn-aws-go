package awsautoscaling


// Experimental.
type AwsAutoscalingPolicy_MetricSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#target_value AwsAutoscalingPolicy#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_capacity_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_capacity_metric_specification AwsAutoscalingPolicy#customized_capacity_metric_specification}
	// Experimental.
	CustomizedCapacityMetricSpecification *AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// customized_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_load_metric_specification AwsAutoscalingPolicy#customized_load_metric_specification}
	// Experimental.
	CustomizedLoadMetricSpecification *AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationProperty `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// customized_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_scaling_metric_specification AwsAutoscalingPolicy#customized_scaling_metric_specification}
	// Experimental.
	CustomizedScalingMetricSpecification *AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationProperty `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// predefined_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_load_metric_specification AwsAutoscalingPolicy#predefined_load_metric_specification}
	// Experimental.
	PredefinedLoadMetricSpecification *AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationProperty `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// predefined_metric_pair_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_metric_pair_specification AwsAutoscalingPolicy#predefined_metric_pair_specification}
	// Experimental.
	PredefinedMetricPairSpecification *AwsAutoscalingPolicy_PredefinedMetricPairSpecificationProperty `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// predefined_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_scaling_metric_specification AwsAutoscalingPolicy#predefined_scaling_metric_specification}
	// Experimental.
	PredefinedScalingMetricSpecification *AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationProperty `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

