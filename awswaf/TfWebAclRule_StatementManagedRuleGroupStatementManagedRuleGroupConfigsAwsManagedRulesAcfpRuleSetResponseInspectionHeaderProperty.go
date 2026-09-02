package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#failure_values TfWebAclRule#failure_values}.
	// Experimental.
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#name TfWebAclRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#success_values TfWebAclRule#success_values}.
	// Experimental.
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

