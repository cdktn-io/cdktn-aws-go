package iotcore


// Experimental.
type AwsTopicRule_IotEventsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#input_name AwsTopicRule#input_name}.
	// Experimental.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn AwsTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#batch_mode AwsTopicRule#batch_mode}.
	// Experimental.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#message_id AwsTopicRule#message_id}.
	// Experimental.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
}

