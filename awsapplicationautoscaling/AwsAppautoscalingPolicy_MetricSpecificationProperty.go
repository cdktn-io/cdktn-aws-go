package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingPolicy_MetricSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#target_value AwsAppautoscalingPolicy#target_value}.
	// Experimental.
	TargetValue *string `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_capacity_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_capacity_metric_specification AwsAppautoscalingPolicy#customized_capacity_metric_specification}
	// Experimental.
	CustomizedCapacityMetricSpecification *AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationProperty `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// customized_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_load_metric_specification AwsAppautoscalingPolicy#customized_load_metric_specification}
	// Experimental.
	CustomizedLoadMetricSpecification *AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// customized_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_scaling_metric_specification AwsAppautoscalingPolicy#customized_scaling_metric_specification}
	// Experimental.
	CustomizedScalingMetricSpecification *AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationProperty `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// predefined_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_load_metric_specification AwsAppautoscalingPolicy#predefined_load_metric_specification}
	// Experimental.
	PredefinedLoadMetricSpecification *AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationProperty `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// predefined_metric_pair_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_metric_pair_specification AwsAppautoscalingPolicy#predefined_metric_pair_specification}
	// Experimental.
	PredefinedMetricPairSpecification *AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationProperty `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// predefined_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_scaling_metric_specification AwsAppautoscalingPolicy#predefined_scaling_metric_specification}
	// Experimental.
	PredefinedScalingMetricSpecification *AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationProperty `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

