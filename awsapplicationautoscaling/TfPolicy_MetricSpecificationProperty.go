package awsapplicationautoscaling


// Experimental.
type TfPolicy_MetricSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#target_value TfPolicy#target_value}.
	// Experimental.
	TargetValue *string `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_capacity_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_capacity_metric_specification TfPolicy#customized_capacity_metric_specification}
	// Experimental.
	CustomizedCapacityMetricSpecification *TfPolicy_CustomizedCapacityMetricSpecificationProperty `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// customized_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_load_metric_specification TfPolicy#customized_load_metric_specification}
	// Experimental.
	CustomizedLoadMetricSpecification *TfPolicy_CustomizedLoadMetricSpecificationProperty `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// customized_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_scaling_metric_specification TfPolicy#customized_scaling_metric_specification}
	// Experimental.
	CustomizedScalingMetricSpecification *TfPolicy_CustomizedScalingMetricSpecificationProperty `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// predefined_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_load_metric_specification TfPolicy#predefined_load_metric_specification}
	// Experimental.
	PredefinedLoadMetricSpecification *TfPolicy_PredefinedLoadMetricSpecificationProperty `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// predefined_metric_pair_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_metric_pair_specification TfPolicy#predefined_metric_pair_specification}
	// Experimental.
	PredefinedMetricPairSpecification *TfPolicy_PredefinedMetricPairSpecificationProperty `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// predefined_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_scaling_metric_specification TfPolicy#predefined_scaling_metric_specification}
	// Experimental.
	PredefinedScalingMetricSpecification *TfPolicy_PredefinedScalingMetricSpecificationProperty `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

