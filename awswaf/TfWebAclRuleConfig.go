package awswaf

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Rule name, unique within the Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#name TfWebAclRule#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Rule priority. Rules with lower priority are evaluated first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#priority TfWebAclRule#priority}
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// ARN of the Web ACL to add the rule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#web_acl_arn TfWebAclRule#web_acl_arn}
	// Experimental.
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#action TfWebAclRule#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#captcha_config TfWebAclRule#captcha_config}
	// Experimental.
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// challenge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#challenge_config TfWebAclRule#challenge_config}
	// Experimental.
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#override_action TfWebAclRule#override_action}
	// Experimental.
	OverrideAction interface{} `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#region TfWebAclRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#rule_label TfWebAclRule#rule_label}
	// Experimental.
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#statement TfWebAclRule#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#timeouts TfWebAclRule#timeouts}
	// Experimental.
	Timeouts *TfWebAclRule_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#visibility_config TfWebAclRule#visibility_config}
	// Experimental.
	VisibilityConfig interface{} `field:"optional" json:"visibilityConfig" yaml:"visibilityConfig"`
}

