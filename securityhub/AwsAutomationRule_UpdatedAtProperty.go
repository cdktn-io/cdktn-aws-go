package securityhub


// Experimental.
type AwsAutomationRule_UpdatedAtProperty struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#date_range AwsAutomationRule#date_range}
	// Experimental.
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#end AwsAutomationRule#end}.
	// Experimental.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#start AwsAutomationRule#start}.
	// Experimental.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

