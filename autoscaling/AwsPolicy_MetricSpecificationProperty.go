package autoscaling


// Experimental.
type AwsPolicy_MetricSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#target_value AwsPolicy#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_capacity_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_capacity_metric_specification AwsPolicy#customized_capacity_metric_specification}
	// Experimental.
	CustomizedCapacityMetricSpecification *AwsPolicy_CustomizedCapacityMetricSpecificationProperty `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// customized_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_load_metric_specification AwsPolicy#customized_load_metric_specification}
	// Experimental.
	CustomizedLoadMetricSpecification *AwsPolicy_CustomizedLoadMetricSpecificationProperty `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// customized_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_scaling_metric_specification AwsPolicy#customized_scaling_metric_specification}
	// Experimental.
	CustomizedScalingMetricSpecification *AwsPolicy_CustomizedScalingMetricSpecificationProperty `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// predefined_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_load_metric_specification AwsPolicy#predefined_load_metric_specification}
	// Experimental.
	PredefinedLoadMetricSpecification *AwsPolicy_PredefinedLoadMetricSpecificationProperty `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// predefined_metric_pair_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_metric_pair_specification AwsPolicy#predefined_metric_pair_specification}
	// Experimental.
	PredefinedMetricPairSpecification *AwsPolicy_PredefinedMetricPairSpecificationProperty `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// predefined_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_scaling_metric_specification AwsPolicy#predefined_scaling_metric_specification}
	// Experimental.
	PredefinedScalingMetricSpecification *AwsPolicy_PredefinedScalingMetricSpecificationProperty `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

