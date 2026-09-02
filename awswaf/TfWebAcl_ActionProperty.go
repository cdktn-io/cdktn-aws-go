package awswaf


// Experimental.
type TfWebAcl_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#allow TfWebAcl#allow}
	// Experimental.
	Allow *TfWebAcl_RuleActionAllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#block TfWebAcl#block}
	// Experimental.
	Block *TfWebAcl_RuleActionBlockProperty `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#captcha TfWebAcl#captcha}
	// Experimental.
	Captcha *TfWebAcl_CaptchaProperty `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#challenge TfWebAcl#challenge}
	// Experimental.
	Challenge *TfWebAcl_ChallengeProperty `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#count TfWebAcl#count}
	// Experimental.
	Count *TfWebAcl_RuleActionCountProperty `field:"optional" json:"count" yaml:"count"`
}

