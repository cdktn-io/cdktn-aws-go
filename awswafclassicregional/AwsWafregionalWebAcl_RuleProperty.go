package awswafclassicregional


// Experimental.
type AwsWafregionalWebAcl_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#priority AwsWafregionalWebAcl#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#rule_id AwsWafregionalWebAcl#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#action AwsWafregionalWebAcl#action}
	// Experimental.
	Action *AwsWafregionalWebAcl_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#override_action AwsWafregionalWebAcl#override_action}
	// Experimental.
	OverrideAction *AwsWafregionalWebAcl_OverrideActionProperty `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#type AwsWafregionalWebAcl#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

