package awsdlm


// Experimental.
type TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#count TfLifecyclePolicy#count}.
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval TfLifecyclePolicy#interval}.
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval_unit TfLifecyclePolicy#interval_unit}.
	// Experimental.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

