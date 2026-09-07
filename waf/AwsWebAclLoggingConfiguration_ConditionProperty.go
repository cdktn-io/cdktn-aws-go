package waf


// Experimental.
type AwsWebAclLoggingConfiguration_ConditionProperty struct {
	// action_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#action_condition AwsWebAclLoggingConfiguration#action_condition}
	// Experimental.
	ActionCondition *AwsWebAclLoggingConfiguration_ActionConditionProperty `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// label_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#label_name_condition AwsWebAclLoggingConfiguration#label_name_condition}
	// Experimental.
	LabelNameCondition *AwsWebAclLoggingConfiguration_LabelNameConditionProperty `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

