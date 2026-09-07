package recyclebin


// Experimental.
type AwsRule_UnlockDelayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#unlock_delay_unit AwsRule#unlock_delay_unit}.
	// Experimental.
	UnlockDelayUnit *string `field:"required" json:"unlockDelayUnit" yaml:"unlockDelayUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#unlock_delay_value AwsRule#unlock_delay_value}.
	// Experimental.
	UnlockDelayValue *float64 `field:"required" json:"unlockDelayValue" yaml:"unlockDelayValue"`
}

