package dlm


// Experimental.
type AwsLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval AwsLifecyclePolicy#interval}.
	// Experimental.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval_unit AwsLifecyclePolicy#interval_unit}.
	// Experimental.
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

