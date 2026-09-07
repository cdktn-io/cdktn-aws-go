package wafclassicregional


// Experimental.
type AwsWebAcl_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#priority AwsWebAcl#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#rule_id AwsWebAcl#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#action AwsWebAcl#action}
	// Experimental.
	Action *AwsWebAcl_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#override_action AwsWebAcl#override_action}
	// Experimental.
	OverrideAction *AwsWebAcl_OverrideActionProperty `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#type AwsWebAcl#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

