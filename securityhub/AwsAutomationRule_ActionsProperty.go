package securityhub


// Experimental.
type AwsAutomationRule_ActionsProperty struct {
	// finding_fields_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#finding_fields_update AwsAutomationRule#finding_fields_update}
	// Experimental.
	FindingFieldsUpdate interface{} `field:"optional" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#type AwsAutomationRule#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

