package awsiotcore


// Experimental.
type TfTopicRule_KinesisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#stream_name TfTopicRule#stream_name}.
	// Experimental.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#partition_key TfTopicRule#partition_key}.
	// Experimental.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

