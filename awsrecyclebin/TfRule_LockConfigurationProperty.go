package awsrecyclebin


// Experimental.
type TfRule_LockConfigurationProperty struct {
	// unlock_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#unlock_delay TfRule#unlock_delay}
	// Experimental.
	UnlockDelay *TfRule_UnlockDelayProperty `field:"required" json:"unlockDelay" yaml:"unlockDelay"`
}

