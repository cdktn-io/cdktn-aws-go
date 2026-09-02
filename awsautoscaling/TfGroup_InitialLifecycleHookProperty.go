package awsautoscaling


// Experimental.
type TfGroup_InitialLifecycleHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#lifecycle_transition TfGroup#lifecycle_transition}.
	// Experimental.
	LifecycleTransition *string `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#name TfGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#default_result TfGroup#default_result}.
	// Experimental.
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#heartbeat_timeout TfGroup#heartbeat_timeout}.
	// Experimental.
	HeartbeatTimeout *float64 `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#notification_metadata TfGroup#notification_metadata}.
	// Experimental.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#notification_target_arn TfGroup#notification_target_arn}.
	// Experimental.
	NotificationTargetArn *string `field:"optional" json:"notificationTargetArn" yaml:"notificationTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#role_arn TfGroup#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

