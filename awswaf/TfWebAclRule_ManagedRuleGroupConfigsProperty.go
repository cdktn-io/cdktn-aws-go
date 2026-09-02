package awswaf


// Experimental.
type TfWebAclRule_ManagedRuleGroupConfigsProperty struct {
	// aws_managed_rules_acfp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_acfp_rule_set TfWebAclRule#aws_managed_rules_acfp_rule_set}
	// Experimental.
	AwsManagedRulesAcfpRuleSet interface{} `field:"optional" json:"awsManagedRulesAcfpRuleSet" yaml:"awsManagedRulesAcfpRuleSet"`
	// aws_managed_rules_anti_ddos_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_anti_ddos_rule_set TfWebAclRule#aws_managed_rules_anti_ddos_rule_set}
	// Experimental.
	AwsManagedRulesAntiDdosRuleSet interface{} `field:"optional" json:"awsManagedRulesAntiDdosRuleSet" yaml:"awsManagedRulesAntiDdosRuleSet"`
	// aws_managed_rules_atp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_atp_rule_set TfWebAclRule#aws_managed_rules_atp_rule_set}
	// Experimental.
	AwsManagedRulesAtpRuleSet interface{} `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
	// aws_managed_rules_bot_control_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_bot_control_rule_set TfWebAclRule#aws_managed_rules_bot_control_rule_set}
	// Experimental.
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#login_path TfWebAclRule#login_path}.
	// Experimental.
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#password_field TfWebAclRule#password_field}
	// Experimental.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#payload_type TfWebAclRule#payload_type}.
	// Experimental.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#username_field TfWebAclRule#username_field}
	// Experimental.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

