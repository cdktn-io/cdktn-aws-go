package awsautoscalingplans


// Experimental.
type TfScalingPlan_PredefinedScalingMetricSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predefined_scaling_metric_type TfScalingPlan#predefined_scaling_metric_type}.
	// Experimental.
	PredefinedScalingMetricType *string `field:"required" json:"predefinedScalingMetricType" yaml:"predefinedScalingMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#resource_label TfScalingPlan#resource_label}.
	// Experimental.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

