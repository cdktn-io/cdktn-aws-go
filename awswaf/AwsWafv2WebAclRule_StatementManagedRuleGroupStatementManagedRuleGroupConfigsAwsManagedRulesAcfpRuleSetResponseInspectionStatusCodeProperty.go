package awswaf


// Experimental.
type AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionStatusCodeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#failure_codes AwsWafv2WebAclRule#failure_codes}.
	// Experimental.
	FailureCodes *[]*float64 `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#success_codes AwsWafv2WebAclRule#success_codes}.
	// Experimental.
	SuccessCodes *[]*float64 `field:"required" json:"successCodes" yaml:"successCodes"`
}

