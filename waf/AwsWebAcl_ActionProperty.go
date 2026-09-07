package waf


// Experimental.
type AwsWebAcl_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#allow AwsWebAcl#allow}
	// Experimental.
	Allow *AwsWebAcl_RuleActionAllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#block AwsWebAcl#block}
	// Experimental.
	Block *AwsWebAcl_RuleActionBlockProperty `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#captcha AwsWebAcl#captcha}
	// Experimental.
	Captcha *AwsWebAcl_CaptchaProperty `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#challenge AwsWebAcl#challenge}
	// Experimental.
	Challenge *AwsWebAcl_ChallengeProperty `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#count AwsWebAcl#count}
	// Experimental.
	Count *AwsWebAcl_RuleActionCountProperty `field:"optional" json:"count" yaml:"count"`
}

