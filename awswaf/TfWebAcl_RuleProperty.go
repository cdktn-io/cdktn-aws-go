package awswaf


// Experimental.
type TfWebAcl_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#name TfWebAcl#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#priority TfWebAcl#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#visibility_config TfWebAcl#visibility_config}
	// Experimental.
	VisibilityConfig *TfWebAcl_RuleVisibilityConfigProperty `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#action TfWebAcl#action}
	// Experimental.
	Action *TfWebAcl_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#captcha_config TfWebAcl#captcha_config}
	// Experimental.
	CaptchaConfig *TfWebAcl_RuleCaptchaConfigProperty `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// challenge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#challenge_config TfWebAcl#challenge_config}
	// Experimental.
	ChallengeConfig *TfWebAcl_RuleChallengeConfigProperty `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#override_action TfWebAcl#override_action}
	// Experimental.
	OverrideAction *TfWebAcl_OverrideActionProperty `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#rule_label TfWebAcl#rule_label}
	// Experimental.
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#statement TfWebAcl#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

