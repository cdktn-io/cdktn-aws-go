package awssecurityhub


// Experimental.
type TfAutomationRule_ActionsProperty struct {
	// finding_fields_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#finding_fields_update TfAutomationRule#finding_fields_update}
	// Experimental.
	FindingFieldsUpdate interface{} `field:"optional" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#type TfAutomationRule#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

