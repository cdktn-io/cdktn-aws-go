package recyclebin


// Experimental.
type AwsRule_LockConfigurationProperty struct {
	// unlock_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#unlock_delay AwsRule#unlock_delay}
	// Experimental.
	UnlockDelay *AwsRule_UnlockDelayProperty `field:"required" json:"unlockDelay" yaml:"unlockDelay"`
}

