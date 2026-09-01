package awswaf


// Experimental.
type AwsWafv2WebAcl_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#allow AwsWafv2WebAcl#allow}
	// Experimental.
	Allow *AwsWafv2WebAcl_RuleActionAllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#block AwsWafv2WebAcl#block}
	// Experimental.
	Block *AwsWafv2WebAcl_RuleActionBlockProperty `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#captcha AwsWafv2WebAcl#captcha}
	// Experimental.
	Captcha *AwsWafv2WebAcl_CaptchaProperty `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#challenge AwsWafv2WebAcl#challenge}
	// Experimental.
	Challenge *AwsWafv2WebAcl_ChallengeProperty `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#count AwsWafv2WebAcl#count}
	// Experimental.
	Count *AwsWafv2WebAcl_RuleActionCountProperty `field:"optional" json:"count" yaml:"count"`
}

