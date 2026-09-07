package securityhub


// Experimental.
type AwsAutomationRuleV2_CriteriaProperty struct {
	// JSON-encoded OCSF finding criteria for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#ocsf_finding_criteria_json AwsAutomationRuleV2#ocsf_finding_criteria_json}
	// Experimental.
	OcsfFindingCriteriaJson *string `field:"required" json:"ocsfFindingCriteriaJson" yaml:"ocsfFindingCriteriaJson"`
}

