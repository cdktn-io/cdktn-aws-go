package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingScheduledAction_ScalableTargetActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_scheduled_action#max_capacity AwsAppautoscalingScheduledAction#max_capacity}.
	// Experimental.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_scheduled_action#min_capacity AwsAppautoscalingScheduledAction#min_capacity}.
	// Experimental.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

