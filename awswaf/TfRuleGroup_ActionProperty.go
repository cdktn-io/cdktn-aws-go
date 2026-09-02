package awswaf


// Experimental.
type TfRuleGroup_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#allow TfRuleGroup#allow}
	// Experimental.
	Allow *TfRuleGroup_AllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#block TfRuleGroup#block}
	// Experimental.
	Block *TfRuleGroup_BlockProperty `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#captcha TfRuleGroup#captcha}
	// Experimental.
	Captcha *TfRuleGroup_CaptchaProperty `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#challenge TfRuleGroup#challenge}
	// Experimental.
	Challenge *TfRuleGroup_ChallengeProperty `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#count TfRuleGroup#count}
	// Experimental.
	Count *TfRuleGroup_CountProperty `field:"optional" json:"count" yaml:"count"`
}

