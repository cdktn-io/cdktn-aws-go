package awsiotcore


// Experimental.
type AwsIotTopicRule_ErrorActionDynamodbv2Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn AwsIotTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// put_item block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#put_item AwsIotTopicRule#put_item}
	// Experimental.
	PutItem *AwsIotTopicRule_ErrorActionDynamodbv2PutItemProperty `field:"optional" json:"putItem" yaml:"putItem"`
}

