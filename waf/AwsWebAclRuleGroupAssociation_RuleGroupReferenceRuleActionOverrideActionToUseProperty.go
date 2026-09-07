package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#allow AwsWebAclRuleGroupAssociation#allow}
	// Experimental.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#block AwsWebAclRuleGroupAssociation#block}
	// Experimental.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#captcha AwsWebAclRuleGroupAssociation#captcha}
	// Experimental.
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#challenge AwsWebAclRuleGroupAssociation#challenge}
	// Experimental.
	Challenge interface{} `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#count AwsWebAclRuleGroupAssociation#count}
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

