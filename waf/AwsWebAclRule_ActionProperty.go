package waf


// Experimental.
type AwsWebAclRule_ActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#allow AwsWebAclRule#allow}
	// Experimental.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#block AwsWebAclRule#block}
	// Experimental.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#captcha AwsWebAclRule#captcha}
	// Experimental.
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#challenge AwsWebAclRule#challenge}
	// Experimental.
	Challenge interface{} `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#count AwsWebAclRule#count}
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

