package iotcore


// Experimental.
type AwsTopicRule_ErrorActionSqsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#queue_url AwsTopicRule#queue_url}.
	// Experimental.
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn AwsTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#use_base64 AwsTopicRule#use_base64}.
	// Experimental.
	UseBase64 interface{} `field:"required" json:"useBase64" yaml:"useBase64"`
}

