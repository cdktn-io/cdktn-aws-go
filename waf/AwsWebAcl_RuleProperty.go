package waf


// Experimental.
type AwsWebAcl_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#name AwsWebAcl#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#priority AwsWebAcl#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#visibility_config AwsWebAcl#visibility_config}
	// Experimental.
	VisibilityConfig *AwsWebAcl_RuleVisibilityConfigProperty `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#action AwsWebAcl#action}
	// Experimental.
	Action *AwsWebAcl_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#captcha_config AwsWebAcl#captcha_config}
	// Experimental.
	CaptchaConfig *AwsWebAcl_RuleCaptchaConfigProperty `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// challenge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#challenge_config AwsWebAcl#challenge_config}
	// Experimental.
	ChallengeConfig *AwsWebAcl_RuleChallengeConfigProperty `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#override_action AwsWebAcl#override_action}
	// Experimental.
	OverrideAction *AwsWebAcl_OverrideActionProperty `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#rule_label AwsWebAcl#rule_label}
	// Experimental.
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#statement AwsWebAcl#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

