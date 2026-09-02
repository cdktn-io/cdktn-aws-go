package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#payload_type TfWebAclRule#payload_type}.
	// Experimental.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#password_field TfWebAclRule#password_field}
	// Experimental.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#username_field TfWebAclRule#username_field}
	// Experimental.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

