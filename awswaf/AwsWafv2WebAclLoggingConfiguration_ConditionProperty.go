package awswaf


// Experimental.
type AwsWafv2WebAclLoggingConfiguration_ConditionProperty struct {
	// action_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#action_condition AwsWafv2WebAclLoggingConfiguration#action_condition}
	// Experimental.
	ActionCondition *AwsWafv2WebAclLoggingConfiguration_ActionConditionProperty `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// label_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#label_name_condition AwsWafv2WebAclLoggingConfiguration#label_name_condition}
	// Experimental.
	LabelNameCondition *AwsWafv2WebAclLoggingConfiguration_LabelNameConditionProperty `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

