package awsiotcore


// Experimental.
type TfTopicRule_Dynamodbv2Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// put_item block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#put_item TfTopicRule#put_item}
	// Experimental.
	PutItem *TfTopicRule_Dynamodbv2PutItemProperty `field:"optional" json:"putItem" yaml:"putItem"`
}

