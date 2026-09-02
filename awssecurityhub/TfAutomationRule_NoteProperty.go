package awssecurityhub


// Experimental.
type TfAutomationRule_NoteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#text TfAutomationRule#text}.
	// Experimental.
	Text *string `field:"required" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#updated_by TfAutomationRule#updated_by}.
	// Experimental.
	UpdatedBy *string `field:"required" json:"updatedBy" yaml:"updatedBy"`
}

