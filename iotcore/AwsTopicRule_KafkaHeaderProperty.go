package iotcore


// Experimental.
type AwsTopicRule_KafkaHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#key AwsTopicRule#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#value AwsTopicRule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

