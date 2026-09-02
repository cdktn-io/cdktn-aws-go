package awssesmailmanager


// Experimental.
type TfRuleSet_PublishToSnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#role_arn TfRuleSet#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#topic_arn TfRuleSet#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy TfRuleSet#action_failure_policy}.
	// Experimental.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#encoding TfRuleSet#encoding}.
	// Experimental.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#payload_type TfRuleSet#payload_type}.
	// Experimental.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
}

