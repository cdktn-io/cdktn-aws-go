package awswaf


// Experimental.
type AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionBodyContainsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#failure_strings AwsWafv2WebAclRule#failure_strings}.
	// Experimental.
	FailureStrings *[]*string `field:"required" json:"failureStrings" yaml:"failureStrings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#success_strings AwsWafv2WebAclRule#success_strings}.
	// Experimental.
	SuccessStrings *[]*string `field:"required" json:"successStrings" yaml:"successStrings"`
}

