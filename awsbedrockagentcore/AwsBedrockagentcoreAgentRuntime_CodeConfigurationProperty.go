package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreAgentRuntime_CodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#entry_point AwsBedrockagentcoreAgentRuntime#entry_point}.
	// Experimental.
	EntryPoint *[]*string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#runtime AwsBedrockagentcoreAgentRuntime#runtime}.
	// Experimental.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#code AwsBedrockagentcoreAgentRuntime#code}
	// Experimental.
	Code interface{} `field:"optional" json:"code" yaml:"code"`
}

