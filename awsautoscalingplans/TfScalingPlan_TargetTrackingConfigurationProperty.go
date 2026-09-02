package awsautoscalingplans


// Experimental.
type TfScalingPlan_TargetTrackingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#target_value TfScalingPlan#target_value}.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#customized_scaling_metric_specification TfScalingPlan#customized_scaling_metric_specification}
	// Experimental.
	CustomizedScalingMetricSpecification *TfScalingPlan_CustomizedScalingMetricSpecificationProperty `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#disable_scale_in TfScalingPlan#disable_scale_in}.
	// Experimental.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#estimated_instance_warmup TfScalingPlan#estimated_instance_warmup}.
	// Experimental.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// predefined_scaling_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predefined_scaling_metric_specification TfScalingPlan#predefined_scaling_metric_specification}
	// Experimental.
	PredefinedScalingMetricSpecification *TfScalingPlan_PredefinedScalingMetricSpecificationProperty `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#scale_in_cooldown TfScalingPlan#scale_in_cooldown}.
	// Experimental.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#scale_out_cooldown TfScalingPlan#scale_out_cooldown}.
	// Experimental.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

