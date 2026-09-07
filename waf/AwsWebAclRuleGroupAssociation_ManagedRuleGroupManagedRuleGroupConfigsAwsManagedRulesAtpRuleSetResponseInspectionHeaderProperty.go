package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionHeaderProperty struct {
	// Strings that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#failure_values AwsWebAclRuleGroupAssociation#failure_values}
	// Experimental.
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// Name of the HTTP header to inspect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#name AwsWebAclRuleGroupAssociation#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Strings that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#success_values AwsWebAclRuleGroupAssociation#success_values}
	// Experimental.
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

