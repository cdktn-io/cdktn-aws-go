package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupProperty struct {
	// Name of the managed rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#name AwsWebAclRuleGroupAssociation#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the managed rule group vendor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#vendor_name AwsWebAclRuleGroupAssociation#vendor_name}
	// Experimental.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// managed_rule_group_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#managed_rule_group_configs AwsWebAclRuleGroupAssociation#managed_rule_group_configs}
	// Experimental.
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#rule_action_override AwsWebAclRuleGroupAssociation#rule_action_override}
	// Experimental.
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
	// Version of the managed rule group. Omit this to use the default version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#version AwsWebAclRuleGroupAssociation#version}
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

