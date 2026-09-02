package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionBodyContainsProperty struct {
	// Strings that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#failure_strings TfWebAclRuleGroupAssociation#failure_strings}
	// Experimental.
	FailureStrings *[]*string `field:"required" json:"failureStrings" yaml:"failureStrings"`
	// Strings that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#success_strings TfWebAclRuleGroupAssociation#success_strings}
	// Experimental.
	SuccessStrings *[]*string `field:"required" json:"successStrings" yaml:"successStrings"`
}

