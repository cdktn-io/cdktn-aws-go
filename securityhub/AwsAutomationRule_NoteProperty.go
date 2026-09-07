package securityhub


// Experimental.
type AwsAutomationRule_NoteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#text AwsAutomationRule#text}.
	// Experimental.
	Text *string `field:"required" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#updated_by AwsAutomationRule#updated_by}.
	// Experimental.
	UpdatedBy *string `field:"required" json:"updatedBy" yaml:"updatedBy"`
}

