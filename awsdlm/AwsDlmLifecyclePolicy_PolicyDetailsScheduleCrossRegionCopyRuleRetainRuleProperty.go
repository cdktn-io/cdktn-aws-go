package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval AwsDlmLifecyclePolicy#interval}.
	// Experimental.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval_unit AwsDlmLifecyclePolicy#interval_unit}.
	// Experimental.
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

