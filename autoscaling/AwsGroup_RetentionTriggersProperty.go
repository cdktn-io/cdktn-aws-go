package autoscaling


// Experimental.
type AwsGroup_RetentionTriggersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#terminate_hook_abandon AwsGroup#terminate_hook_abandon}.
	// Experimental.
	TerminateHookAbandon *string `field:"optional" json:"terminateHookAbandon" yaml:"terminateHookAbandon"`
}

