package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#payload_type TfWebAclRule#payload_type}.
	// Experimental.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// address_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#address_fields TfWebAclRule#address_fields}
	// Experimental.
	AddressFields interface{} `field:"optional" json:"addressFields" yaml:"addressFields"`
	// email_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#email_field TfWebAclRule#email_field}
	// Experimental.
	EmailField interface{} `field:"optional" json:"emailField" yaml:"emailField"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#password_field TfWebAclRule#password_field}
	// Experimental.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// phone_number_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#phone_number_fields TfWebAclRule#phone_number_fields}
	// Experimental.
	PhoneNumberFields interface{} `field:"optional" json:"phoneNumberFields" yaml:"phoneNumberFields"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#username_field TfWebAclRule#username_field}
	// Experimental.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

