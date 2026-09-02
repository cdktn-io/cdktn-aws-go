package awsautoscaling


// Experimental.
type TfPolicy_TargetTrackingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#target_value TfPolicy#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_metric_specification TfPolicy#customized_metric_specification}
	// Experimental.
	CustomizedMetricSpecification *TfPolicy_CustomizedMetricSpecificationProperty `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#disable_scale_in TfPolicy#disable_scale_in}.
	// Experimental.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_metric_specification TfPolicy#predefined_metric_specification}
	// Experimental.
	PredefinedMetricSpecification *TfPolicy_PredefinedMetricSpecificationProperty `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
}

