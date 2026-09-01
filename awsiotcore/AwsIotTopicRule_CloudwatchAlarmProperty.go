package awsiotcore


// Experimental.
type AwsIotTopicRule_CloudwatchAlarmProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#alarm_name AwsIotTopicRule#alarm_name}.
	// Experimental.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn AwsIotTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#state_reason AwsIotTopicRule#state_reason}.
	// Experimental.
	StateReason *string `field:"required" json:"stateReason" yaml:"stateReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#state_value AwsIotTopicRule#state_value}.
	// Experimental.
	StateValue *string `field:"required" json:"stateValue" yaml:"stateValue"`
}

