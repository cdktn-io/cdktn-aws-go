package awsautoscaling


// Experimental.
type TfGroup_InstanceLifecyclePolicyProperty struct {
	// retention_triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#retention_triggers TfGroup#retention_triggers}
	// Experimental.
	RetentionTriggers *TfGroup_RetentionTriggersProperty `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}

