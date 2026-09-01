package awsiotcore


// Experimental.
type AwsIotTopicRule_ErrorActionRepublishProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn AwsIotTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#topic AwsIotTopicRule#topic}.
	// Experimental.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#qos AwsIotTopicRule#qos}.
	// Experimental.
	Qos *float64 `field:"optional" json:"qos" yaml:"qos"`
}

