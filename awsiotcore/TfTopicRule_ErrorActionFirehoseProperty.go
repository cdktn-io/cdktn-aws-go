package awsiotcore


// Experimental.
type TfTopicRule_ErrorActionFirehoseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#delivery_stream_name TfTopicRule#delivery_stream_name}.
	// Experimental.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#batch_mode TfTopicRule#batch_mode}.
	// Experimental.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#separator TfTopicRule#separator}.
	// Experimental.
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

