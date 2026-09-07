package applicationautoscaling


// Experimental.
type AwsPolicy_TargetTrackingScalingPolicyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#target_value AwsPolicy#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_metric_specification AwsPolicy#customized_metric_specification}
	// Experimental.
	CustomizedMetricSpecification *AwsPolicy_CustomizedMetricSpecificationProperty `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#disable_scale_in AwsPolicy#disable_scale_in}.
	// Experimental.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_metric_specification AwsPolicy#predefined_metric_specification}
	// Experimental.
	PredefinedMetricSpecification *AwsPolicy_PredefinedMetricSpecificationProperty `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scale_in_cooldown AwsPolicy#scale_in_cooldown}.
	// Experimental.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scale_out_cooldown AwsPolicy#scale_out_cooldown}.
	// Experimental.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

