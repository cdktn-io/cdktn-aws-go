package awsiotcore


// Experimental.
type TfTopicRule_ErrorActionKafkaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#client_properties TfTopicRule#client_properties}.
	// Experimental.
	ClientProperties *map[string]*string `field:"required" json:"clientProperties" yaml:"clientProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#destination_arn TfTopicRule#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#topic TfTopicRule#topic}.
	// Experimental.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#header TfTopicRule#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#key TfTopicRule#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#partition TfTopicRule#partition}.
	// Experimental.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
}

