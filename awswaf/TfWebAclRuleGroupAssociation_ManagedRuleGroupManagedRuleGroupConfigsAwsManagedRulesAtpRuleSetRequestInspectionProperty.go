package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspectionProperty struct {
	// Payload type for inspection, either JSON or FORM_ENCODED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#payload_type TfWebAclRuleGroupAssociation#payload_type}
	// Experimental.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#password_field TfWebAclRuleGroupAssociation#password_field}
	// Experimental.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#username_field TfWebAclRuleGroupAssociation#username_field}
	// Experimental.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

