package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_ShareRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target_accounts AwsDlmLifecyclePolicy#target_accounts}.
	// Experimental.
	TargetAccounts *[]*string `field:"required" json:"targetAccounts" yaml:"targetAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#unshare_interval AwsDlmLifecyclePolicy#unshare_interval}.
	// Experimental.
	UnshareInterval *float64 `field:"optional" json:"unshareInterval" yaml:"unshareInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#unshare_interval_unit AwsDlmLifecyclePolicy#unshare_interval_unit}.
	// Experimental.
	UnshareIntervalUnit *string `field:"optional" json:"unshareIntervalUnit" yaml:"unshareIntervalUnit"`
}

