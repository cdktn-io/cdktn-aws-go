package awsiotcore


// Experimental.
type TfTopicRule_IotAnalyticsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#channel_name TfTopicRule#channel_name}.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#batch_mode TfTopicRule#batch_mode}.
	// Experimental.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
}

