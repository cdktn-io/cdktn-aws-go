package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingTarget_SuspendedStateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_target#dynamic_scaling_in_suspended AwsAppautoscalingTarget#dynamic_scaling_in_suspended}.
	// Experimental.
	DynamicScalingInSuspended interface{} `field:"optional" json:"dynamicScalingInSuspended" yaml:"dynamicScalingInSuspended"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_target#dynamic_scaling_out_suspended AwsAppautoscalingTarget#dynamic_scaling_out_suspended}.
	// Experimental.
	DynamicScalingOutSuspended interface{} `field:"optional" json:"dynamicScalingOutSuspended" yaml:"dynamicScalingOutSuspended"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_target#scheduled_scaling_suspended AwsAppautoscalingTarget#scheduled_scaling_suspended}.
	// Experimental.
	ScheduledScalingSuspended interface{} `field:"optional" json:"scheduledScalingSuspended" yaml:"scheduledScalingSuspended"`
}

