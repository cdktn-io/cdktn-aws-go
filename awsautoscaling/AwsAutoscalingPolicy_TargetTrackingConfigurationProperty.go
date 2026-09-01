package awsautoscaling


// Experimental.
type AwsAutoscalingPolicy_TargetTrackingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#target_value AwsAutoscalingPolicy#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#customized_metric_specification AwsAutoscalingPolicy#customized_metric_specification}
	// Experimental.
	CustomizedMetricSpecification *AwsAutoscalingPolicy_CustomizedMetricSpecificationProperty `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#disable_scale_in AwsAutoscalingPolicy#disable_scale_in}.
	// Experimental.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#predefined_metric_specification AwsAutoscalingPolicy#predefined_metric_specification}
	// Experimental.
	PredefinedMetricSpecification *AwsAutoscalingPolicy_PredefinedMetricSpecificationProperty `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
}

