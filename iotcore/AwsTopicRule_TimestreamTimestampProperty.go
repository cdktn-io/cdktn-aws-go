package iotcore


// Experimental.
type AwsTopicRule_TimestreamTimestampProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#unit AwsTopicRule#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#value AwsTopicRule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

