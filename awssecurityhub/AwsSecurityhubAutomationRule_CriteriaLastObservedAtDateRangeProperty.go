package awssecurityhub


// Experimental.
type AwsSecurityhubAutomationRule_CriteriaLastObservedAtDateRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#unit AwsSecurityhubAutomationRule#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#value AwsSecurityhubAutomationRule#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

