package autoscaling


// Experimental.
type AwsGroup_InstanceLifecyclePolicyProperty struct {
	// retention_triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#retention_triggers AwsGroup#retention_triggers}
	// Experimental.
	RetentionTriggers *AwsGroup_RetentionTriggersProperty `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}

