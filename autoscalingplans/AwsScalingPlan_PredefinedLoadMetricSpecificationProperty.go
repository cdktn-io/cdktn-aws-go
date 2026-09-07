package autoscalingplans


// Experimental.
type AwsScalingPlan_PredefinedLoadMetricSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predefined_load_metric_type AwsScalingPlan#predefined_load_metric_type}.
	// Experimental.
	PredefinedLoadMetricType *string `field:"required" json:"predefinedLoadMetricType" yaml:"predefinedLoadMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#resource_label AwsScalingPlan#resource_label}.
	// Experimental.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

