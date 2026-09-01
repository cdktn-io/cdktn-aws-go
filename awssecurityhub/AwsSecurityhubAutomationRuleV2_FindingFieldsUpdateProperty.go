package awssecurityhub


// Experimental.
type AwsSecurityhubAutomationRuleV2_FindingFieldsUpdateProperty struct {
	// A comment for the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#comment AwsSecurityhubAutomationRuleV2#comment}
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The severity ID to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#severity_id AwsSecurityhubAutomationRuleV2#severity_id}
	// Experimental.
	SeverityId *float64 `field:"optional" json:"severityId" yaml:"severityId"`
	// The status ID to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#status_id AwsSecurityhubAutomationRuleV2#status_id}
	// Experimental.
	StatusId *float64 `field:"optional" json:"statusId" yaml:"statusId"`
}

