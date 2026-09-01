package awsbedrockagents


// Experimental.
type AwsBedrockagentPrompt_ToolConfigurationProperty struct {
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#tool AwsBedrockagentPrompt#tool}
	// Experimental.
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
	// tool_choice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#tool_choice AwsBedrockagentPrompt#tool_choice}
	// Experimental.
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
}

