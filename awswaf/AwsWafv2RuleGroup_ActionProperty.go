package awswaf


// Experimental.
type AwsWafv2RuleGroup_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#allow AwsWafv2RuleGroup#allow}
	// Experimental.
	Allow *AwsWafv2RuleGroup_AllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#block AwsWafv2RuleGroup#block}
	// Experimental.
	Block *AwsWafv2RuleGroup_BlockProperty `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#captcha AwsWafv2RuleGroup#captcha}
	// Experimental.
	Captcha *AwsWafv2RuleGroup_CaptchaProperty `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#challenge AwsWafv2RuleGroup#challenge}
	// Experimental.
	Challenge *AwsWafv2RuleGroup_ChallengeProperty `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#count AwsWafv2RuleGroup#count}
	// Experimental.
	Count *AwsWafv2RuleGroup_CountProperty `field:"optional" json:"count" yaml:"count"`
}

