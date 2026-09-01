package awsbedrockagents


// Experimental.
type AwsBedrockagentPrompt_ToolChoiceProperty struct {
	// any block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#any AwsBedrockagentPrompt#any}
	// Experimental.
	Any interface{} `field:"optional" json:"any" yaml:"any"`
	// auto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#auto AwsBedrockagentPrompt#auto}
	// Experimental.
	Auto interface{} `field:"optional" json:"auto" yaml:"auto"`
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#tool AwsBedrockagentPrompt#tool}
	// Experimental.
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
}

