package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingPolicy_TargetTrackingScalingPolicyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#target_value AwsAppautoscalingPolicy#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#customized_metric_specification AwsAppautoscalingPolicy#customized_metric_specification}
	// Experimental.
	CustomizedMetricSpecification *AwsAppautoscalingPolicy_CustomizedMetricSpecificationProperty `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#disable_scale_in AwsAppautoscalingPolicy#disable_scale_in}.
	// Experimental.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predefined_metric_specification AwsAppautoscalingPolicy#predefined_metric_specification}
	// Experimental.
	PredefinedMetricSpecification *AwsAppautoscalingPolicy_PredefinedMetricSpecificationProperty `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scale_in_cooldown AwsAppautoscalingPolicy#scale_in_cooldown}.
	// Experimental.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scale_out_cooldown AwsAppautoscalingPolicy#scale_out_cooldown}.
	// Experimental.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

