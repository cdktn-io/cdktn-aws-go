package awsbedrockagents


// Experimental.
type AwsBedrockagentAgentActionGroup_ApiSchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#payload AwsBedrockagentAgentActionGroup#payload}.
	// Experimental.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#s3 AwsBedrockagentAgentActionGroup#s3}
	// Experimental.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

