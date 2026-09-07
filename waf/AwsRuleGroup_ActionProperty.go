package waf


// Experimental.
type AwsRuleGroup_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#allow AwsRuleGroup#allow}
	// Experimental.
	Allow *AwsRuleGroup_AllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#block AwsRuleGroup#block}
	// Experimental.
	Block *AwsRuleGroup_BlockProperty `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#captcha AwsRuleGroup#captcha}
	// Experimental.
	Captcha *AwsRuleGroup_CaptchaProperty `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#challenge AwsRuleGroup#challenge}
	// Experimental.
	Challenge *AwsRuleGroup_ChallengeProperty `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#count AwsRuleGroup#count}
	// Experimental.
	Count *AwsRuleGroup_CountProperty `field:"optional" json:"count" yaml:"count"`
}

