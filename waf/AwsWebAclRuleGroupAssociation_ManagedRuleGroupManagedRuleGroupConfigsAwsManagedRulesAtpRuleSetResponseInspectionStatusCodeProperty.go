package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCodeProperty struct {
	// Status codes that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#failure_codes AwsWebAclRuleGroupAssociation#failure_codes}
	// Experimental.
	FailureCodes *[]*float64 `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Status codes that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#success_codes AwsWebAclRuleGroupAssociation#success_codes}
	// Experimental.
	SuccessCodes *[]*float64 `field:"required" json:"successCodes" yaml:"successCodes"`
}

