package awssecurityhub


// Experimental.
type AwsSecurityhubAutomationRuleV2_ActionProperty struct {
	// The action type: FINDING_FIELDS_UPDATE or EXTERNAL_INTEGRATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#type AwsSecurityhubAutomationRuleV2#type}
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// external_integration_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#external_integration_configuration AwsSecurityhubAutomationRuleV2#external_integration_configuration}
	// Experimental.
	ExternalIntegrationConfiguration interface{} `field:"optional" json:"externalIntegrationConfiguration" yaml:"externalIntegrationConfiguration"`
	// finding_fields_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#finding_fields_update AwsSecurityhubAutomationRuleV2#finding_fields_update}
	// Experimental.
	FindingFieldsUpdate interface{} `field:"optional" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
}

