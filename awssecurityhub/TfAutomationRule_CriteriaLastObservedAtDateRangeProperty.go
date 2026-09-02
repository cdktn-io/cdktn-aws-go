package awssecurityhub


// Experimental.
type TfAutomationRule_CriteriaLastObservedAtDateRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#unit TfAutomationRule#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#value TfAutomationRule#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

