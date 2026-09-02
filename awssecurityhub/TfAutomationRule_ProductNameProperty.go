package awssecurityhub


// Experimental.
type TfAutomationRule_ProductNameProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#comparison TfAutomationRule#comparison}.
	// Experimental.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#value TfAutomationRule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

