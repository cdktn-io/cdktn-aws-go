package awsiotcore


// Experimental.
type TfTopicRule_ErrorActionDynamodbProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#hash_key_field TfTopicRule#hash_key_field}.
	// Experimental.
	HashKeyField *string `field:"required" json:"hashKeyField" yaml:"hashKeyField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#hash_key_value TfTopicRule#hash_key_value}.
	// Experimental.
	HashKeyValue *string `field:"required" json:"hashKeyValue" yaml:"hashKeyValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#table_name TfTopicRule#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#hash_key_type TfTopicRule#hash_key_type}.
	// Experimental.
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#operation TfTopicRule#operation}.
	// Experimental.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#payload_field TfTopicRule#payload_field}.
	// Experimental.
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#range_key_field TfTopicRule#range_key_field}.
	// Experimental.
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#range_key_type TfTopicRule#range_key_type}.
	// Experimental.
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#range_key_value TfTopicRule#range_key_value}.
	// Experimental.
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
}

