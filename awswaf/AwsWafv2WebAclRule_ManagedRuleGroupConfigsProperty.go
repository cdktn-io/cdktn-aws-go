package awswaf


// Experimental.
type AwsWafv2WebAclRule_ManagedRuleGroupConfigsProperty struct {
	// aws_managed_rules_acfp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_acfp_rule_set AwsWafv2WebAclRule#aws_managed_rules_acfp_rule_set}
	// Experimental.
	AwsManagedRulesAcfpRuleSet interface{} `field:"optional" json:"awsManagedRulesAcfpRuleSet" yaml:"awsManagedRulesAcfpRuleSet"`
	// aws_managed_rules_anti_ddos_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_anti_ddos_rule_set AwsWafv2WebAclRule#aws_managed_rules_anti_ddos_rule_set}
	// Experimental.
	AwsManagedRulesAntiDdosRuleSet interface{} `field:"optional" json:"awsManagedRulesAntiDdosRuleSet" yaml:"awsManagedRulesAntiDdosRuleSet"`
	// aws_managed_rules_atp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_atp_rule_set AwsWafv2WebAclRule#aws_managed_rules_atp_rule_set}
	// Experimental.
	AwsManagedRulesAtpRuleSet interface{} `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
	// aws_managed_rules_bot_control_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_bot_control_rule_set AwsWafv2WebAclRule#aws_managed_rules_bot_control_rule_set}
	// Experimental.
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#login_path AwsWafv2WebAclRule#login_path}.
	// Experimental.
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#password_field AwsWafv2WebAclRule#password_field}
	// Experimental.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#payload_type AwsWafv2WebAclRule#payload_type}.
	// Experimental.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#username_field AwsWafv2WebAclRule#username_field}
	// Experimental.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

