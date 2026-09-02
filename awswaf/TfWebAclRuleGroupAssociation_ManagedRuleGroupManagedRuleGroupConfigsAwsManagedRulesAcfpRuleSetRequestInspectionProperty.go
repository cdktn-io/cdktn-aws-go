package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionProperty struct {
	// Payload type for inspection, either JSON or FORM_ENCODED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#payload_type TfWebAclRuleGroupAssociation#payload_type}
	// Experimental.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// address_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#address_fields TfWebAclRuleGroupAssociation#address_fields}
	// Experimental.
	AddressFields interface{} `field:"optional" json:"addressFields" yaml:"addressFields"`
	// email_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#email_field TfWebAclRuleGroupAssociation#email_field}
	// Experimental.
	EmailField interface{} `field:"optional" json:"emailField" yaml:"emailField"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#password_field TfWebAclRuleGroupAssociation#password_field}
	// Experimental.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// phone_number_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#phone_number_fields TfWebAclRuleGroupAssociation#phone_number_fields}
	// Experimental.
	PhoneNumberFields interface{} `field:"optional" json:"phoneNumberFields" yaml:"phoneNumberFields"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#username_field TfWebAclRuleGroupAssociation#username_field}
	// Experimental.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

