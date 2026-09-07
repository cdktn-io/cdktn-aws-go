package bedrockagentcore


// Experimental.
type AwsAgentRuntime_CodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#entry_point AwsAgentRuntime#entry_point}.
	// Experimental.
	EntryPoint *[]*string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#runtime AwsAgentRuntime#runtime}.
	// Experimental.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#code AwsAgentRuntime#code}
	// Experimental.
	Code interface{} `field:"optional" json:"code" yaml:"code"`
}

