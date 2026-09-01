package awssecurityhub


// Experimental.
type AwsSecurityhubAutomationRule_NoteUpdatedAtProperty struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#date_range AwsSecurityhubAutomationRule#date_range}
	// Experimental.
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#end AwsSecurityhubAutomationRule#end}.
	// Experimental.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#start AwsSecurityhubAutomationRule#start}.
	// Experimental.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

