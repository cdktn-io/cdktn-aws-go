package awsiotcore


// Experimental.
type AwsIotTopicRuleDestination_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule_destination#create AwsIotTopicRuleDestination#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule_destination#delete AwsIotTopicRuleDestination#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule_destination#update AwsIotTopicRuleDestination#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

