package awssecurityhub


// Experimental.
type TfAutomationRule_FirstObservedAtProperty struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#date_range TfAutomationRule#date_range}
	// Experimental.
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#end TfAutomationRule#end}.
	// Experimental.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#start TfAutomationRule#start}.
	// Experimental.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

