package bedrockagents


// Experimental.
type AwsFlow_ToolConfigurationProperty struct {
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#tool AwsFlow#tool}
	// Experimental.
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
	// tool_choice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#tool_choice AwsFlow#tool_choice}
	// Experimental.
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
}

