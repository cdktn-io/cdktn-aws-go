package awsiotcore


// Experimental.
type TfTopicRule_StepFunctionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#state_machine_name TfTopicRule#state_machine_name}.
	// Experimental.
	StateMachineName *string `field:"required" json:"stateMachineName" yaml:"stateMachineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#execution_name_prefix TfTopicRule#execution_name_prefix}.
	// Experimental.
	ExecutionNamePrefix *string `field:"optional" json:"executionNamePrefix" yaml:"executionNamePrefix"`
}

