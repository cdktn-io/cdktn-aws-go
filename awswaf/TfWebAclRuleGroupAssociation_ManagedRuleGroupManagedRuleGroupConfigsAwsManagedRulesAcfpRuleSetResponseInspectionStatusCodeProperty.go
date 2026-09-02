package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionStatusCodeProperty struct {
	// Status codes that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#failure_codes TfWebAclRuleGroupAssociation#failure_codes}
	// Experimental.
	FailureCodes *[]*float64 `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Status codes that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#success_codes TfWebAclRuleGroupAssociation#success_codes}
	// Experimental.
	SuccessCodes *[]*float64 `field:"required" json:"successCodes" yaml:"successCodes"`
}

