package awsapplicationautoscaling


// Experimental.
type TfPolicy_PredefinedMetricPairSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_metric_type TfPolicy#predefined_metric_type}.
	// Experimental.
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#resource_label TfPolicy#resource_label}.
	// Experimental.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

