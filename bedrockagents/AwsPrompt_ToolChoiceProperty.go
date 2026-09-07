package bedrockagents


// Experimental.
type AwsPrompt_ToolChoiceProperty struct {
	// any block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#any AwsPrompt#any}
	// Experimental.
	Any interface{} `field:"optional" json:"any" yaml:"any"`
	// auto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#auto AwsPrompt#auto}
	// Experimental.
	Auto interface{} `field:"optional" json:"auto" yaml:"auto"`
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#tool AwsPrompt#tool}
	// Experimental.
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
}

