package waf

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRuleGroupAssociationConfig struct {
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
	// Priority of the rule within the Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#priority AwsWebAclRuleGroupAssociation#priority}
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Name of the rule to create in the Web ACL that references the rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#rule_name AwsWebAclRuleGroupAssociation#rule_name}
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// ARN of the Web ACL to associate the Rule Group with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#web_acl_arn AwsWebAclRuleGroupAssociation#web_acl_arn}
	// Experimental.
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
	// managed_rule_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#managed_rule_group AwsWebAclRuleGroupAssociation#managed_rule_group}
	// Experimental.
	ManagedRuleGroup interface{} `field:"optional" json:"managedRuleGroup" yaml:"managedRuleGroup"`
	// Override action for the rule group. Valid values are 'none' and 'count'. Defaults to 'none'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#override_action AwsWebAclRuleGroupAssociation#override_action}
	// Experimental.
	OverrideAction *string `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#region AwsWebAclRuleGroupAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule_group_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#rule_group_reference AwsWebAclRuleGroupAssociation#rule_group_reference}
	// Experimental.
	RuleGroupReference interface{} `field:"optional" json:"ruleGroupReference" yaml:"ruleGroupReference"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#timeouts AwsWebAclRuleGroupAssociation#timeouts}
	// Experimental.
	Timeouts *AwsWebAclRuleGroupAssociation_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#visibility_config AwsWebAclRuleGroupAssociation#visibility_config}
	// Experimental.
	VisibilityConfig interface{} `field:"optional" json:"visibilityConfig" yaml:"visibilityConfig"`
}

